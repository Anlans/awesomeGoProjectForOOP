package main

import (
	"awesomeProject2/model"
	"fmt"
)

func main() {
	xm := model.NewUserInfo("I'm XiaoMing")

	fmt.Println(xm.Name)

	product := model.Product{}
	product.SetProductName("apple")
	fmt.Println(product.GetProductName())

	alipay := &model.Alipay{
		PaymentArgs:  model.PaymentArgs{
			AppID: "alipay123",
			MchID: "alipaymcid",
			Key: "alipaysdfsdfsdfsd",
			CallbackUrl: "https://api.imooc.com/alipay",
		},
		AlipayOpenID: "alipayopenid",
	}

	weixinpay := &model.WeixinPay{
		PaymentArgs:  model.PaymentArgs{
			AppID: "weixin123",
			MchID: "weixinmcid",
			Key: "weixinsdfsdfsdfsd",
			CallbackUrl: "https://api.imooc.com/weixinpay",
		},
		WeixinOpenID: "weixinpayopenid",
	}

	alipay.PaymentOther.AppID = "alipay_paymentother"
	fmt.Println(alipay.PaymentOther.AppID)
	fmt.Println(weixinpay.PaymentArgs.AppID)

	paymentArgs := model.PaymentArgs{
		AppID:       "superAppid",
		MchID:       "superMchid",
		Key:         "superKey",
		CallbackUrl: "superCallbackurl",
	}

	paymentArgs.Info()
	alipay.Info()
	weixinpay.Info()

	//  多态
	_payment := &payment{paymentmethod: "alipay"}
	_payment.info()
	_payment.topay()
	/*
	type pay interface {  //定义 pay 接口，里面规定两种方法,topay 和 info
		topay()
		info()
	}
	type payment struct {
		paymentmethod string

	}
	func (this *payment)topay() {  //对 pay 接口的方法 topay 和 info 进行实现
		fmt.Println("topay: ", this.paymentmethod)
	}

	func (this *payment)info() {
		fmt.Println("info: ", this.paymentmethod)
	}
	*/
	var _pay pay
	/*type pay interface {
		topay()
		info()
	}*/
	_pay = _payment
	_pay.info()
	_pay.topay()

	var _readwrite readwrite
	_readwrite.echo()

	var _Integer Integer
	_Integer.echo()
	fmt.Println("------------------------------------------------")
	//多态的最佳实践
	log := &Log{
		name: "微信小程序支付日志",
		content: "微信小程序支付日志内容",
		addtime: 0,
	}
	var _iowrite *iowrite
	var _netwrite *netwrite
	log.writeLog(_iowrite)
	log.writeLog(_netwrite)

	/*
	>>>:
	微信小程序支付日志---微信小程序支付日志内容
	iowrite:echo()
	iowrite:out()
	微信小程序支付日志---微信小程序支付日志内容
	netwrite:echo()
	netwrite:out()
	*/

	netlog := &NetLog{Log{
		name: "微信小程序支付日志，网络",
		content: "微信小程序支付日志内容，网络",
		addtime: 0,
	}}
	netlog.writeLog(_netwrite)

	/*
	>>>:
	微信小程序支付日志，网络---微信小程序支付日志内容，网络
	netwrite:echo()
	netwrite:out()
	*/

	filelog := &IOLog{Log{
		name:    "微信小程序支付日志,文件",
		content: "微信小程序支付日志内容,文件",
		addtime: 0,
	}}
	filelog.writeLog(_iowrite)
	/*
	>>>:
	微信小程序支付日志,文件---微信小程序支付日志内容,文件
	iowrite:echo()
	iowrite:out()
	*/
}

//日志结构体
type Log struct {
	name string
	content string
	addtime int64
}

//给Log增加方法
func (this *Log)writeLog(_write write) {
	fmt.Println(this.name+"---"+this.content)
	_write.echo()
	_write.out()
}

type  NetLog struct {
	Log		//继承自Log结构体
}

type IOLog struct {
	Log
}


type iowrite string
//让这两个类型：iowrite,netwrite.分别实现write接口
func (this *iowrite)echo() {
	fmt.Println("iowrite:echo()")
}
func (this *iowrite)out() {
	fmt.Println("iowrite:out()")
}

type netwrite string
func(this *netwrite)echo() {
	fmt.Println("netwrite:echo()")
}
func(this *netwrite)out() {
	fmt.Println("netwrite:out()")
}


type readerwriter interface {//这句话为何还在，是否可去
	write
	read
}
//定义读、写接口
type write interface {
	echo()//这两个方法:echo,out之后用来被其他玩意实现
	out()
}

type read interface {
	scan()
	input()
}
//自定义类型，基于内置类型 string
type readwrite string

func (this *readwrite)echo() {
	fmt.Println("readWrite:echo()")
}

func (this *readwrite)out() {
	fmt.Println("readWrite:out()")
}

func (this *readwrite)scan() {
	fmt.Println("readWrite:scan()")
}

func (this *readwrite)input() {
	fmt.Println("readWrite:input()")
}


//自定义类型，基于内置类型 int
type Integer int

func (this *Integer)echo() {
	fmt.Println("Integer: echo()")
}

func (this *Integer)out() {
	fmt.Println("Integer: out()")
}

func (this *Integer)scan() {
	fmt.Println("Integer: scan()")
}

func (this *Integer)input() {
	fmt.Println("Integer: input()")
}


type pay interface {
	topay()
	info()
}

type payment struct {
	paymentmethod string

}
//给payment增加方法，topay和info
func (this *payment)topay() {
	fmt.Println("topay: ", this.paymentmethod)
}

func (this *payment)info() {
	fmt.Println("info: ", this.paymentmethod)
}