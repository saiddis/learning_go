package main

type Order struct {
	id       int
	product  string
	quantity int
	price    float64
}

type OrderManager struct {
	products map[int]string
	quantity map[string]int
	prices   map[int]float64
}

func (om OrderManager) AddOrder(id int, product string, quantity int, price float64) {
	order := Order{
		id:       id,
		product:  product,
		quantity: quantity,
		price:    price,
	}

	om.prices[order.id] = order.price
	om.products[order.id] = order.product
	om.quantity[order.product] = order.quantity
}

func (om OrderManager) ReturnOrder(id, quantity int) {
	for i, v := range om.products {
		if i == id {
			om.quantity[v] -= quantity
		}
	}
}

func (om OrderManager) TotalActiveOrders() float64 {
	var total float64

	for i, v := range om.products {
		total += om.prices[i] * float64(om.quantity[v])
	}

	return total
}

func NewOrderManager() OrderManager {
	return OrderManager{
		map[int]string{},
		map[string]int{},
		map[int]float64{},
	}
}
