package go_skrill

import (
	"encoding/json"
	"github.com/asaka1234/go-skrill/utils"
	"strconv"
	"strings"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req SkrillDepositBackReq, processor func(SkrillDepositBackReq) error) error {
	//验证签名
	// Verify MD5 signature
	key := cli.getDepositBackMD5(req)
	if req.PayToEmail != "demoqco@sun-fish.com" && key != req.Md5sig {
		reqJSON, _ := json.Marshal(req)
		cli.logger.Errorf("Skrill#depositBack#verify,req:%s,key:%s", string(reqJSON), key)
		return nil
	}
	
	//开始处理
	return processor(req)
}

func (cli *Client) getDepositBackMD5(cbRsp SkrillDepositBackReq) string {
	// Concatenate all the required fields
	data := cli.DepositId + cbRsp.TransactionID + strings.ToUpper(cli.DepositSetting) +
		cbRsp.MbAmount + cbRsp.MbCurrency + strconv.Itoa(cbRsp.Status)

	// Create MD5 hash
	return utils.GetMD5([]byte(data))
}
