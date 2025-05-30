package go_skrill

import (
	"github.com/asaka1234/go-skrill/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params SkrillInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params SkrillInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}
