package model

import "fmt"

type Alipay struct {
	PaymentArgs
	PaymentOther
	AlipayOpenID string
}

func (this *Alipay)Info() {
	fmt.Printf("alipayInfo = %v\n", this)
}