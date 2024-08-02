package main

import "fmt"

type Transaction struct {
	transactionType string
	amount          float64
}

type Account struct {
	balance            float64
	transactionHistory []Transaction
}

func (a *Account) AddTransaction(transType string, amount float64) {
	if a.balance < amount {
		fmt.Println("Not enough balane for this transaction")
		return
	}

	trans := Transaction{
		transactionType: transType,
		amount:          amount,
	}

	a.balance -= amount
	a.transactionHistory = append(a.transactionHistory, trans)
}

func (a Account) Balance() float64 {
	return a.balance
}
