package main

type BankAccount interface {
	Deposit(amount float64)
	WithDraw(amount float64)
}

type Account struct {
	balance *float64
}

func (a Account) Deposit(amount float64) {
	*a.balance += amount
}

func (a Account) WithDraw(amount float64) {
	*a.balance -= amount
}

func (a Account) GetBalance() float64 {
	return *a.balance
}

type AccountHandler struct {
	BankAccount
}

func NewAccountHandler(ba BankAccount) AccountHandler {
	return AccountHandler{
		ba,
	}
}
