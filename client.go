package go_skrill

import (
	"github.com/asaka1234/go-skrill/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	DepositMerchantId  string //商户号
	DepositEmail       string //从用户的srill账号扣款,划转给cpt的账号.  所以这里是cpt的账号
	DepositSetting     string //webhook返回的数据,需要做md5签名验证, 是做这个用的.
	DepositUrl         string
	DepositCallbackUrl string //充值回调

	WithdrawMerchantId       string //商户号
	WithdrawMerchantEmail    string //psp分配的账号
	WithdrawMerchantPassword string //psp分配的密码
	WithdrawUrl              string //充值回调

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, depositId, depositEmail, depositSetting, depositUrl, depositCallbackUrl, withdrawId, withdrawMerchantEmail, withdrawMerchantPassword, withdrawUrl string) *Client {
	return &Client{
		DepositMerchantId:        depositId, //商户号
		DepositEmail:             depositEmail,
		DepositSetting:           depositSetting,
		DepositUrl:               depositUrl,
		DepositCallbackUrl:       depositCallbackUrl, //充值回调
		WithdrawMerchantId:       withdrawId,         //提现回调
		WithdrawMerchantEmail:    withdrawMerchantEmail,
		WithdrawMerchantPassword: withdrawMerchantPassword,
		WithdrawUrl:              withdrawUrl, //充值回调

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
