package go_skrill

import (
	"github.com/fatih/structs"
)

func (cli *Client) Deposit(req SkrillDepositReq) (map[string]interface{}, error) {

	paramMap := structs.Map(req)
	return paramMap, nil
}
