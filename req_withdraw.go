package go_skrill

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/asaka1234/go-skrill/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/cast"
)

//文档： https://www.skrill.com/fileadmin/content/pdf/Skrill_Automated_Payments_Interface_Guide.pdf 的4章节:4. SEND MONEY USING AN HTTPS REQUEST

// 整体提现分2步
// 1. 准备一个session环境
// 2. 在session环境中执行transfer  (Sending a transfer prepare request)
// 每一步的返回都是一个xml
func (cli *Client) Withdraw(req SkrillWithdrawReq) (*SkrillWithdrawResponse, error) {
	// 1. 准备一个session环境
	sid, err := cli.InitSession(req)
	if err != nil {
		return nil, err
	}
	// 2. 在session中pre-order
	return cli.SendWithdrawRequest(sid)
}

// -------第一步:准备一个session环境----------------
// GET https://www.skrill.com/app/pay.pl?action=prepare&email=merchant@host.com&password=6b4c1ba48880bcd3341dbaeb68b2647f&amount=1.2&currency=EUR&bnf_email=beneficiary@domain.com&subject=some_subject&note=some_note&frn_trn_id=111
/*
	//成功
	<?xml version="1.0" encoding="UTF-8"?>
	<response>
		<sid>5e281d1376d92ba789ca7f0583e045d4</sid>
	</response>

	//失败
	<?xml version="1.0" encoding="UTF-8"?>
	<response>
		<error>
			<error_msg>MISSING_AMOUNT</error_msg>
		</error>
	</response>
*/
func (cli *Client) InitSession(req SkrillWithdrawReq) (string, error) {
	rawURL := cli.Params.WithdrawUrl

	// 1. 准备一个session环境, 返回的是这个session的id
	params := map[string]string{
		"action":     "prepare",
		"email":      cli.Params.WithdrawMerchantEmail,    //Your merchant account email address
		"password":   cli.Params.WithdrawMerchantPassword, //Your MD5 API/MQI password.
		"amount":     cast.ToString(req.PayAmount),
		"currency":   cast.ToString(req.PayCurrency),      //EUR
		"bnf_email":  req.UserEmail,                       //收到钱的人的邮箱
		"subject":    "Withdraw:" + cast.ToString(req.ID), //Subject of the notification email. 发给邮箱的邮件内容
		"note":       "Withdraw:" + cast.ToString(req.ID), //Comment to be included in the notification email.  邮件里的一个备注
		"frn_trn_id": cast.ToString(req.ID),               //唯一id
	}

	//-----------------------------------------------------
	// 预下单
	// Step 3: Send prepare request
	// 发送HTTP请求
	resp1, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetFormData(params).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp1))
	cli.logger.Infof("PSPResty#skrill#init->%+v", string(restLog))

	if err != nil {
		cli.logger.Errorf("请求失败: %s", err.Error())
		return "", err
	}

	if resp1.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return "", fmt.Errorf("status code: %d", resp1.StatusCode())
	}

	if resp1.Error() != nil {
		//反序列化错误会在此捕捉
		return "", fmt.Errorf("%s", resp1.Error())
	}

	body := resp1.Body()

	// Step 4: Parse XML response
	type Error struct {
		ErrorMsg string `xml:"error_msg"`
	}

	type Response struct {
		XMLName xml.Name `xml:"response"`
		Error   Error    `xml:"error"`
		SID     string   `xml:"sid"`
	}

	var resp Response
	if err := xml.Unmarshal(body, &resp); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	if resp.Error.ErrorMsg != "" {
		return "", fmt.Errorf("skrill error: %s", resp.Error.ErrorMsg)
	}

	sid := resp.SID
	//cli.logger.Infof("skrill withdraw sid %s", sid)

	return sid, nil
}

// GET https://www.skrill.com/app/pay.pl?action=transfer&sid=5e281d1376d92ba789ca7f0583e045d4
/*
	<?xml version="1.0" encoding="UTF-8"?>
	<response>
		<transaction>
		<amount>1.20</amount>
		<currency>EUR</currency>
		<id>497029</id>
		<status>2</status>
		<status_msg>processed</status_msg>
		</transaction>
	</response>
*/
// 在session环境中直接来pre-order
func (cli *Client) SendWithdrawRequest(sid string) (*SkrillWithdrawResponse, error) {
	rawURL := cli.Params.WithdrawUrl

	transferParams := map[string]string{
		"action": "transfer",
		"sid":    sid,
	}

	// 发送HTTP请求
	resp1, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetHeaders(getHeaders()).
		SetFormData(transferParams).
		SetDebug(cli.debugMode).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp1))
	cli.logger.Infof("PSPResty#skrill#withdraw->%+v", string(restLog))

	if err != nil {
		cli.logger.Errorf("请求失败: %s", err.Error())
		return nil, err
	}

	if resp1.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp1.StatusCode())
	}

	if resp1.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%s", resp1.Error())
	}

	transferBody := resp1.Body()

	//-----------------------------------------------------

	// Step 6: Parse transfer response
	type Transaction struct {
		ID       string `xml:"id"`     //Transaction ID.
		Status   string `xml:"status"` //枚举: 1 – scheduled (if beneficiary is not yet registered at Skrill),2 - processed (if beneficiary is registered)
		Amount   string `xml:"amount"`
		Currency string `xml:"currency"`
	}

	type Error struct {
		ErrorMsg string `xml:"error_msg"`
	}

	type TransferResponse struct {
		XMLName      xml.Name      `xml:"response"`
		Error        Error         `xml:"error"`
		Transactions []Transaction `xml:"transaction"`
	}

	var transferResp TransferResponse
	if err := xml.Unmarshal(transferBody, &transferResp); err != nil {
		return nil, fmt.Errorf("failed to parse transfer response: %v", err)
	}

	if transferResp.Error.ErrorMsg != "" {
		return nil, fmt.Errorf("skrill transfer error: %s", transferResp.Error.ErrorMsg)
	}

	if len(transferResp.Transactions) == 0 {
		return nil, fmt.Errorf("no transaction in response")
	}

	//拿到psp订单信息
	txn := transferResp.Transactions[0]

	return &SkrillWithdrawResponse{
		ID:       txn.ID,
		Amount:   txn.Amount,
		Currency: txn.Currency,
		Status:   cast.ToInt(txn.Status),
	}, nil
}
