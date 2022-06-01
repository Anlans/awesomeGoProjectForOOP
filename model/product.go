package model

type Product struct {
	productName  string
	productPrice float32
}

//  getter setter
func (this *Product) SetProductName(_productName string) {
	this.productName = _productName
}

func (this *Product)GetProductName() string {
	return this.productName
}

func (this *Product) SetProductPrice(_productPrice float32) {
	this.productPrice = _productPrice
}

func (this *Product)GetProductPrice() float32 {
	return this.productPrice
}