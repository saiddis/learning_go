package main

import "fmt"

type Account2 struct {
	balance float64
	history []string
}

func (a *Account2) Deposit(amount float64) {
	a.history = append(a.history, fmt.Sprintf("deposit: %f", amount))
	a.balance += amount
}

func (a *Account2) Withdraw(amount float64) bool {
	if a.balance < amount {
		return false
	}

	a.history = append(a.history, fmt.Sprintf("withdraw: %f", amount))
	a.balance -= amount
	return true
}

func (a Account2) TransactionHistory() []string {
	return a.history
}
