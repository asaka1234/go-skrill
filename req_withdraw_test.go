package go_skrill

import (
	"fmt"
	"testing"
)

//--------------------------------------------

func TestWithdraw(t *testing.T) {

	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, SkrillInitParams{DepositId, DepositEmail, DepositSetting, DepositUrl, DepositCallbackUrl, WithdrawId, WithdrawMerchantEmail, WithdrawMerchantPassword, WithdrawUrl})

	//发请求
	resp, err := cli.Withdraw(GenWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
}

func GenWithdrawRequestDemo() SkrillWithdrawReq {

	return SkrillWithdrawReq{
		UserID:      1234,
		UserEmail:   "demo@gmail.com",
		PayAmount:   1.00, //outNo
		PayCurrency: "THB",
		ID:          8081081,
	}
}
