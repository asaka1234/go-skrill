package go_skrill

import (
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

//--------------------------------------------

func TestDeposit(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &SkrillInitParams{DepositId, DepositEmail, DepositSetting, DepositUrl, DepositCallbackUrl, SecretWord, WithdrawId, WithdrawMerchantEmail, WithdrawMerchantPassword, WithdrawUrl})

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func TestDepositCallback(t *testing.T) {
	vLog := VLog{}

	cli := NewClient(vLog, &SkrillInitParams{DepositId, DepositEmail, DepositSetting, DepositUrl, DepositCallbackUrl, SecretWord, WithdrawId, WithdrawMerchantEmail, WithdrawMerchantPassword, WithdrawUrl})

	//发请求
	cli.DepositCallback(GetDepositCallbackReq(), DepositBackProcessor)

}

func GetDepositCallbackReq() SkrillDepositBackReq {
	return SkrillDepositBackReq{
		TransactionID: "202507241121330681",
		MbAmount:      "200",
		Amount:        "200",
		Md5sig:        "F18698B327ED616ACD4DA7BEB7643C32",
		MerchantID:    "210526825",
		//PaymentType: "",
		MbTransactionID: "6393550374",
		MbCurrency:      "USD",
		PayFromEmail:    "hamp_31%40yahoo.com",
		PayToEmail:      "fin%40cptinternational.com",
		Currency:        "USD",
		CustomerID:      "168399354",
		Status:          2,
	}
}

func DepositBackProcessor(req SkrillDepositBackReq) error {
	return nil
}

func GenDepositRequestDemo() SkrillDepositReq {

	return SkrillDepositReq{
		Currency:             "MYR",
		Detail1Text:          "220099",     //uid
		Detail1Description:   "1609032335", //outNo
		Amount:               "1.00",
		TransactionID:        "123", //必须要是存在的,不然报错
		Language:             "EN",
		RecipientDescription: "aaa",
	}
}
