package go_skrill

// ----------pre order-------------------------
type SkrillDepositReq struct {
	PayToEmail           string `json:"pay_to_email"`
	Currency             string `json:"currency"`
	Amount               string `json:"amount"`
	Detail1Text          string `json:"detail1_text"`
	Detail1Description   string `json:"detail1_description"`
	StatusURL            string `json:"status_url"`
	Language             string `json:"language"`
	URL                  string `json:"url"` //回调地址?
	PayLink              string `json:"payLink"`
	TransactionID        string `json:"transaction_id"`
	RecipientDescription string `json:"recipient_description"`

	// ReturnURL        string `json:"return_url"`
	// ReturnURLTarget  string `json:"return_url_target"`
}

// ---------------callback-----------------------
type SkrillDepositBackReq struct {
	PayToEmail       string `json:"pay_to_email"`
	PayFromEmail     string `json:"pay_from_email"`
	MerchantID       string `json:"merchant_id"`
	CustomerID       string `json:"customer_id"`
	TransactionID    string `json:"transaction_id"`
	MbTransactionID  string `json:"mb_transaction_id"`
	MbAmount         string `json:"mb_amount"`
	MbCurrency       string `json:"mb_currency"`
	Status           int    `json:"status"` // -2:failed, 2:processed, 0:pending, -1:cancelled, -3:chargeback
	FailedReasonCode string `json:"failed_reason_code"`
	Md5sig           string `json:"md5sig"`
	Sha2sig          string `json:"sha2sig"`
	Amount           string `json:"amount"`
	Currency         string `json:"currency"`
	MerchantFields   string `json:"merchant_fields"`
}

// -----------------------------------
type SkrillWithdrawReq struct {
	UserID      int64   `json:"userId"`      // 用户ID
	PayAmount   float64 `json:"payAmount"`   // 支付金额
	PayCurrency string  `json:"payCurrency"` // 支付货币
	ID          int64   `json:"id"`          //业务订单id
}

//------------------------------------------------------------

type SkrillWithdrawResponse struct {
	ID       string `json:"id"`
	Status   int    `json:"status"`
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
