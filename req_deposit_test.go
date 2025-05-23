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
	cli := NewClient(vLog, DepositId, DepositEmail, DepositSetting, DepositUrl, DepositCallbackUrl, WithdrawId, WithdrawMerchantEmail, WithdrawMerchantPassword, WithdrawUrl)

	//发请求
	resp, err := cli.Deposit(GenDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
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
