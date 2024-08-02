package main

type Sale struct {
	item   string
	amount float64
}

type SalesReport struct {
	Sales []Sale
}

func (sr *SalesReport) AddSale(item string, amount float64) {
	sale := Sale{
		item:   item,
		amount: amount,
	}

	sr.Sales = append(sr.Sales, sale)
}

func (sr SalesReport) TotalSales() float64 {
	var total float64

	for _, v := range sr.Sales {
		total += v.amount
	}

	return total
}
