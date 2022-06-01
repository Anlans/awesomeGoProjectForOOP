package model

import "fmt"

type PaymentArgs struct {
	AppID string
	MchID string
	Key string
	CallbackUrl string
}

func (this *PaymentArgs)Info() {
	fmt.Printf("info = %v\n", this)
}