package go_mpay

import (
	"time"
)

// ----------pre order-------------------------
type MPayDepositReq struct {
	UserID      int64      `json:"userId"`      // 用户ID
	PayAmount   float64    `json:"payAmount"`   // 支付金额
	ID          int64      `json:"id"`          //业务订单id
	PayUrl      string     `json:"payUrl"`      // 支付链接
	PayCurrency string     `json:"payCurrency"` // 支付货币
	UserName    string     `json:"userName"`    // 用户name
	CreateTime  *time.Time `json:"createTime"`  //业务订单撞见时间
}

//------------------------------------------------------------

type MPayDepositResponse struct {
	Code float64 `json:"res"` // 1 是正确
	Msg  string  `json:"msg"`
	Ec   float64 `json:"ec"`
	Data struct {
		Fmt string `json:"fmt"`
		Dtm string `json:"dtm"`
		Lst string `json:"lst"`
	} `json:"dt"`
}

// ---------------callback-----------------------
type MPayDepositBackReq struct {
	OrderId     string `json:"orderId"`
	Amount      string `json:"amount"`
	OrderAmount string `json:"orderAmount"`
	Currency    string `json:"currency"`
	TradeTime   string `json:"tradeTime"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	// 算法：md5
	// 内容：amount=200&currency=CNY&orderId=ES89760987
	//&status=1&tradeTime=1541488344884&type=AlipayChannel
	Sign string `json:"sign"`
	Uuid string `json:"uuid"`
	Key  string `json:"key"`
	// 算法：md5
	// 内容：amount=200&currency=CNY&orderId=ES89760987
	//&status=0&tradeTime=1541488344884&type=AlipayChannel&key=RSA公钥+uuid
	Signature string `json:"signature"`
	// 算法：md5
	// 内容：amount=200&currency=CNY&exAmount=2&exRate=7.0000000000&orderId=ES89760987&sourceCurrency=USD&status=0&tradeTime=1541488344884&type=AlipayChannel&key=RSA公钥+uuid
	NSignature string `json:"_signature"`
	TxID       string `json:"txid"`
	Chain      string `json:"chain"`
	BizId      string `json:"bizId"`
	CustNo     string `json:"custNo"`
	PayAddress string `json:"payAddress"`
	ToAddress  string `json:"ToAddress"`
}

//成功的话,返回 "{\"success\"}"
