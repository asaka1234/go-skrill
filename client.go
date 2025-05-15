package go_skrill

import (
	"github.com/asaka1234/go-skrill/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	DepositId          string //商户号
	DepositEmail       string
	DepositSetting     string
	DepositUrl         string
	DepositCallbackUrl string //充值回调

	WithdrawId               string //提现回调
	WithdrawMerchantEmail    string
	WithdrawMerchantPassword string
	WithdrawUrl              string //充值回调

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, depositId, depositEmail, depositSetting, depositUrl, depositCallbackUrl, withdrawId, withdrawMerchantEmail, withdrawMerchantPassword, withdrawUrl string) *Client {
	return &Client{
		DepositId:                depositId, //商户号
		DepositEmail:             depositEmail,
		DepositSetting:           depositSetting,
		DepositUrl:               depositUrl,
		DepositCallbackUrl:       depositCallbackUrl, //充值回调
		WithdrawId:               withdrawId,         //提现回调
		WithdrawMerchantEmail:    withdrawMerchantEmail,
		WithdrawMerchantPassword: withdrawMerchantPassword,
		WithdrawUrl:              withdrawUrl, //充值回调

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
