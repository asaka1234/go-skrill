package go_skrill

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-skrill/utils"
	"strconv"
	"strings"
)

// https://www.skrill.com/fileadmin/content/pdf/Skrill_Quick_Checkout_Guide.pdf
// 2.5. Skrill status response
func (cli *Client) DepositCallback(req SkrillDepositBackReq, processor func(SkrillDepositBackReq) error) error {
	//验证签名
	expectedSign := cli.getDepositBackMD5(req)
	if expectedSign != req.Md5sig {
		fmt.Println("签名验证失败")
		return errors.New("sign verify failed!")
	}

	//开始处理
	return processor(req)
}

// 章节 9.4. MD5 signature
func (cli *Client) getDepositBackMD5(cbRsp SkrillDepositBackReq) string {
	// Concatenate all the required fields
	data := cli.DepositMerchantId + cbRsp.TransactionID + strings.ToUpper(cli.DepositSetting) +
		cbRsp.MbAmount + cbRsp.MbCurrency + strconv.Itoa(cbRsp.Status)

	// Create MD5 hash
	return utils.GetMD5([]byte(data))
}
