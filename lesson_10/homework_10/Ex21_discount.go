package main

type Discount struct {
	persentage float64
}

func (o *Order) ApplyDiscount(discount Discount) {
	o.price = o.price / float64(100) * discount.persentage
}

func (o Order) TotalAmount() float64 {
	return o.price * float64(o.quantity)
}
