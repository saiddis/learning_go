package main

type Budget interface {
	AddIncome(amount float64)
	AddExpense(amount float64)
}

type MonthlyBudget struct {
	balance *float64
}

func (mb MonthlyBudget) AddIncome(amount float64) {
	*mb.balance += amount
}

func (mb MonthlyBudget) AddExpense(amount float64) {
	*mb.balance -= amount
}

func (mb MonthlyBudget) GetBalance() float64 {
	return *mb.balance
}

type BudgetHandler struct {
	Budget
}

func NewBudgetHandler(b Budget) BudgetHandler {
	return BudgetHandler{
		b,
	}
}
