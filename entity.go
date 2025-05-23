package go_skrill

// ----------pre order-------------------------
type SkrillDepositReq struct {
	Currency             string `json:"currency" mapstructure:"currency"`
	Amount               string `json:"amount" mapstructure:"amount"`
	Detail1Text          string `json:"detail1_text" mapstructure:"detail1_text"`
	Detail1Description   string `json:"detail1_description" mapstructure:"detail1_description"`
	Language             string `json:"language" mapstructure:"language"` //EN
	TransactionID        string `json:"transaction_id" mapstructure:"transaction_id"`
	RecipientDescription string `json:"recipient_description" mapstructure:"recipient_description"` //Your Company Name
	//这个是cpt这个skrill商户的账号, 说明是用这个账号收钱的. 这里sdk设置
	//PayToEmail string `json:"pay_to_email"`
	//这个是实际form提交的地址,也就是skrill提供的一个地址，也是sdk设置
	//URL string `json:"url"`
	//PayLink            string `json:"payLink"` //请求psp的地址
	//这个是回调地址,让sdk设置
	//StatusURL string `json:"status_url" mapstructure:"status_url"` //callback回调地址 . If you have provided a status_url value in your HTML form, Skrill will post the transaction details of each payment to that URL.
}

// ---------------callback-----------------------
type SkrillDepositBackReq struct {
	PayToEmail      string `json:"pay_to_email"`      //收款账号,就是cpt的公司账号
	PayFromEmail    string `json:"pay_from_email"`    //付款账号, 是客户的email账号
	MerchantID      string `json:"merchant_id"`       //Unique ID of your Skrill account.用来计算md5签名的
	MbTransactionID string `json:"mb_transaction_id"` //psp内部的订单号
	MbAmount        string `json:"mb_amount"`
	MbCurrency      string `json:"mb_currency"`
	Status          int    `json:"status"` // -2:failed, 2:processed, 0:pending, -1:cancelled, -3:chargeback
	Md5sig          string `json:"md5sig"` //md5签名, 这里要用来做验签!!!
	Amount          string `json:"amount"`
	Currency        string `json:"currency"`
	//option
	CustomerID       string `json:"customer_id"`
	Sha2sig          string `json:"sha2sig"` //SHA2签名
	MerchantFields   string `json:"merchant_fields"`
	TransactionID    string `json:"transaction_id"` //merchant的订单号
	FailedReasonCode string `json:"failed_reason_code"`
}

// -----------------------------------
type SkrillWithdrawReq struct {
	UserID      int64   `json:"userId"`      // 用户ID
	UserEmail   string  `json:"userEmail"`   // 邮箱
	PayAmount   float64 `json:"payAmount"`   // 支付金额
	PayCurrency string  `json:"payCurrency"` // 支付货币
	ID          int64   `json:"id"`          // 业务订单id
}

//------------------------------------------------------------

type SkrillWithdrawResponse struct {
	ID       string `json:"id"`     //是psp自己的订单号
	Status   int    `json:"status"` ////枚举: 1 – scheduled (if beneficiary is not yet registered at Skrill),2 - processed (if beneficiary is registered)
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
