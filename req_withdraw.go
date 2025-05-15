package go_skrill

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/cast"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func (cli *Client) Withdraw(req SkrillWithdrawReq) (*SkrillWithdrawResponse, error) {
	// Step 1: Get user
	// Step 2: Prepare request
	reqParams := url.Values{}
	reqParams.Set("action", "prepare")
	reqParams.Set("email", cli.WithdrawMerchantEmail)
	reqParams.Set("password", cli.WithdrawMerchantPassword)
	reqParams.Set("amount", cast.ToString(req.PayAmount))
	reqParams.Set("currency", cast.ToString(req.PayCurrency))
	reqParams.Set("bnf_email", req.UserEmail)
	reqParams.Set("subject", "Withdraw:"+cast.ToString(req.ID))
	reqParams.Set("note", "Withdraw:"+cast.ToString(req.ID))
	reqParams.Set("frn_trn_id", cast.ToString(req.ID))

	//-----------------------------------------------------
	// 预下单
	// Step 3: Send prepare request
	prepareRsp, err := http.Get(cli.WithdrawUrl + "?" + reqParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("prepare request failed: %v", err)
	}
	defer prepareRsp.Body.Close()

	body, err := ioutil.ReadAll(prepareRsp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	cli.logger.Infof("skrill withdraw prepareRsp %s", string(body))

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
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	if resp.Error.ErrorMsg != "" {
		return nil, fmt.Errorf("skrill error: %s", resp.Error.ErrorMsg)
	}

	sid := resp.SID
	cli.logger.Infof("skrill withdraw sid %s", sid)

	//-----------------------------------------------------
	// Step 5: Make transfer request
	// 又发了一个请求
	transferParams := url.Values{}
	transferParams.Set("action", "transfer")
	transferParams.Set("sid", sid)

	transferRsp, err := http.Get(cli.WithdrawUrl + "?" + transferParams.Encode())
	if err != nil {
		return nil, fmt.Errorf("transfer request failed: %v", err)
	}
	defer transferRsp.Body.Close()

	transferBody, err := ioutil.ReadAll(transferRsp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read transfer response: %v", err)
	}

	cli.logger.Infof("skrill withdraw transferRsp %s", string(transferBody))

	//-----------------------------------------------------

	// Step 6: Parse transfer response
	type Transaction struct {
		ID       string `xml:"id"`
		Status   string `xml:"status"`
		Amount   string `xml:"amount"`
		Currency string `xml:"currency"`
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

	txn := transferResp.Transactions[0]
	if txn.Status == "1" {
		cli.logger.Infof("beneficiary(%s) is not yet registered at Skrill", req.UserEmail)
	}

	status, err := strconv.Atoi(txn.Status)
	if err != nil {
		return nil, fmt.Errorf("invalid status: %v", err)
	}

	return &SkrillWithdrawResponse{
		ID:       txn.ID,
		Amount:   txn.Amount,
		Currency: txn.Currency,
		Status:   status,
	}, nil
}
