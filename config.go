package go_skrill

const (
	DepositId          = "11111" //商户号
	DepositEmail       = "fin@111.com"
	DepositSetting     = "1111 LIMITED"
	DepositUrl         = "https://pay.skrill.com"
	DepositCallbackUrl = "" //充值回调

	//https://www.skrill.com/fileadmin/content/pdf/Skrill_Automated_Payments_Interface_Guide.pdf
	WithdrawId               = "12222"                             //提现回调
	WithdrawMerchantEmail    = "fin@12222.com"                     //skrill的账号
	WithdrawMerchantPassword = "3edcc"                             //skrill的密码
	WithdrawUrl              = "https://www.skrill.com/app/pay.pl" //提现
)
