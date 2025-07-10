package go_skrill

import (
	"github.com/mitchellh/mapstructure"
)

// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
// 2.3.3. Parameters to be posted to Quick Checkout
func (cli *Client) Deposit(req SkrillDepositReq) (map[string]interface{}, error) {

	var paramMap map[string]interface{}
	mapstructure.Decode(req, &paramMap)

	//最终是让前端构造一个form表单,用以提交请求给到deposit url上去

	//补充公共字段
	paramMap["pay_to_email"] = cli.Params.DepositEmail //给cpt的skrill账户充值
	paramMap["url"] = cli.Params.DepositUrl            //发送请求的psp地址
	paramMap["payLink"] = cli.Params.DepositUrl        //发送请求的psp地址
	paramMap["status_url"] = cli.Params.DepositBackUrl // 回调地址

	return paramMap, nil
}
