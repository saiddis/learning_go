package main

type TaxCalculator interface {
	CalculateIncomeTax(amount float64) float64
	CalculateVAT(amount float64) float64
}

type SimpleTaxCalculator struct {
	taxPercent float64
	vatPercent float64
}

func (stc SimpleTaxCalculator) CalculateIncomeTax(amount float64) float64 {
	return amount / 100 * stc.taxPercent
}

func (stc SimpleTaxCalculator) CalculateVAT(amount float64) float64 {
	return amount / 100 * stc.vatPercent
}

type TaxHandler struct {
	TaxCalculator
}

func NewTaxHandler(tc TaxCalculator) TaxHandler {
	return TaxHandler{
		tc,
	}
}
