package main

type CurrencyConverter interface {
	ToUSD(amount float64) float64
	ToEUR(amount float64) float64
}

type Exchange struct {
	coefToUSD float64
	coefToEUR float64
}

func (e Exchange) ToUSD(amount float64) float64 {
	return e.coefToUSD * amount
}

func (e Exchange) ToEUR(amount float64) float64 {
	return e.coefToEUR * amount
}

type MoneyExchanger struct {
	CurrencyConverter
}

func NewMoneyExchanger(cc CurrencyConverter) MoneyExchanger {
	return MoneyExchanger{
		cc,
	}
}
