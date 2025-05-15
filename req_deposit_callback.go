package go_skrill

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-skrill/utils"
	"strings"
)

// 充值的回调处理(传入一个处理函数)
func (cli *Client) DepositCallback(req SkrillDepositBackReq, processor func(SkrillDepositBackReq) error) error {
	//验证签名

	cli.logger.Infof("mpay<----%+v", req)
	// amount=200&currency=CNY&orderId=ES89760987&status=0&tradeTime=1541488344884&type=AlipayChannel&key=RSA公钥+uuid
	raw := fmt.Sprintf(
		"amount=%s&currency=%s&orderId=%s&status=%s&tradeTime=%s&type=%s&key=%s%s",
		req.Amount,
		req.Currency,
		req.OrderId,
		req.Status,
		req.TradeTime,
		req.Type,
		cli.Secret,
		req.Uuid,
	)
	cli.logger.Infof("mpay<----raw: %s", raw)

	sign := utils.GetMD5([]byte(raw))
	cli.logger.Infof("mpay<----md5: %s == %s ?", sign, req.Signature)

	if strings.ToLower(sign) != strings.ToLower(req.Signature) {
		cli.logger.Errorf("sign verify failed")
		return errors.New("sign verify failed")
	}

	//开始处理
	return processor(req)
}
