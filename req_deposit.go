package go_skrill

import (
	"github.com/mitchellh/mapstructure"
)

// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
// 6.5.1. Example of a Skrill 1‐Tap payment form
func (cli *Client) Deposit(req SkrillDepositReq) (map[string]interface{}, error) {

	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)

	//补充公共字段
	paramMap["pay_to_email"] = cli.DepositEmail     //给cpt的skrill账户充值
	paramMap["url"] = cli.DepositUrl                //发送请求的psp地址
	paramMap["payLink"] = cli.DepositUrl            //发送请求的psp地址
	paramMap["status_url"] = cli.DepositCallbackUrl //发送请求的psp地址

	return paramMap, nil
}
