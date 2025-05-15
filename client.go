package go_mpay

import (
	"github.com/asaka1234/go-mpay/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	BizId       string
	Secret      string
	BaseURL     string
	CallbackURL string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, bizId string, secret string, baseURL string, callbackUrl string) *Client {
	return &Client{
		BizId:  bizId,
		Secret: secret,

		BaseURL:     baseURL,
		CallbackURL: callbackUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
