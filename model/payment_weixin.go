package model

import "fmt"

type WeixinPay struct {
	PaymentArgs	//继承了payment.go里的支付结构体
	WeixinOpenID string
}

func (this *WeixinPay)Info() {
	fmt.Printf("weixinInfo = %v\n", this)
}