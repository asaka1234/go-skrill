package go_skrill

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-skrill/utils"
	"github.com/spf13/cast"
	"time"
)

func (cli *Client) Withdraw(req SkrillWithdrawReq) (*SkrillWithdrawResponse, error) {

	loc := time.FixedZone("UTC+8", 8*60*60)

	// 构建支付URL
	payUrl := fmt.Sprintf("bizId=%s&custNo=%d&amount=%s&orderNo=%d"+
		"&noticeUrl=%s&backUrl=%s&clientIp=%s&currency=%s"+
		"&orderTime=%d&devType=%s&ext=%s",
		cli.BizId,
		req.UserID,
		cast.ToString(req.PayAmount),
		req.ID,
		cli.CallbackURL,
		req.PayUrl,
		utils.GetIP(),
		req.PayCurrency,
		req.CreateTime.In(loc).UnixMilli(),
		"1",
		toJSONString(Ext{
			Bank: Bank{
				UserName: req.UserName,
			},
		}),
	)

	// RSA加密
	encryption, err := utils.EncryptPublicKey(payUrl, cli.Secret)
	if err != nil {
		cli.logger.Errorf("加密失败: %s", err.Error())
		return nil, err
	}

	fmt.Printf("--->raw: %s", payUrl)
	fmt.Printf("--->encryption: %s", encryption)

	//返回值会放到这里
	var result SkrillWithdrawResponse

	// 构建请求参数
	paramMap := map[string]interface{}{
		"bizId":      cli.BizId,
		"encryption": encryption,
		// "paymentTypeShow": "by_predict",
	}

	// 发送HTTP请求
	_, err = cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(paramMap).
		SetResult(&result).
		Post(cli.BaseURL)

	if err != nil {
		cli.logger.Errorf("请求失败: %s", err.Error())
		return nil, err
	}

	fmt.Printf("<---rsp: %+v", result)

	return &result, err
}
