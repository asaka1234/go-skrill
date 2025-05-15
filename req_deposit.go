package go_skrill

import (
	"github.com/fatih/structs"
)

func (cli *Client) Deposit(req SkrillDepositReq) (map[string]interface{}, error) {

	paramMap := structs.Map(req)
	paramMap["pay_to_email"] = cli.DepositEmail
	paramMap["detail1_description"] = "Account:"
	paramMap["status_url"] = cli.DepositCallbackUrl
	paramMap["url"] = cli.DepositCallbackUrl
	paramMap["status_url"] = cli.DepositUrl
	paramMap["payLink"] = cli.DepositUrl
	return paramMap, nil
}
