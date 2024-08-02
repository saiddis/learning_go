package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("---Ex_30---")
	var userActTracker UserActivityTracker
	homeworkTime := time.Now()
	userActTracker.AddActivity("Doing homework", homeworkTime)
	userActTracker.AddActivity("Going to work", time.Now())
	userActTracker.AddActivity("Sleeping", time.Now())
	fmt.Println(userActTracker.ActivityAfterTime(homeworkTime))

	fmt.Println("---Ex_29---")
	var lib Library
	lib.AddBook("Mastery", "Robert Greene", 2013)
	lib.AddBook("Learning Go", "Jon Bodner", 2021)
	lib.AddBook("Deep Learning", "Goodfellow", 2018)
	fmt.Println(lib.BooksByAuthorAfterYear(2015))

	fmt.Println("---Ex_28---")
	var cm ContractManager
	cm.AddContract(213, "someone", 1241234)
	cm.AddContract(244, "someone else", 2342354)
	cm.AddContract(224, "someone", 235345)
	fmt.Println(cm.TotalAmountForClient("someone"))

	fmt.Println("---Ex_27---")
	var um UserManager
	um.AddUser("said", "said@said.com", 17)
	um.AddUser("erpugh", "erp@said.com", 20)
	um.AddUser("wret", "wret@said.com", 24)
	fmt.Println(um.UsersOlderThan(18))

	fmt.Println("---Ex_26---")
	var p Project
	p.AddTask("Safty", "Ensure safty", "Stated")
	p.AddTask("Location", "Alocate new location", "On progress")
	p.AddTask("Money", "Get a loan", "Stated")
	fmt.Println(p.CountTasksByStatus("Stated"))

	fmt.Println("---Ex_25---")
	om := NewOrderManager()
	om.AddOrder(123, "Milk", 1000, 8)
	om.AddOrder(123, "Egg", 3000, 1)
	om.AddOrder(124, "Fish", 500, 50)
	om.ReturnOrder(123, 500)
	fmt.Println(om.TotalActiveOrders())

	fmt.Println("---Ex_24---")
	a := Account{
		balance:            4125,
		transactionHistory: []Transaction{},
	}
	a.AddTransaction("credit", 3425)
	fmt.Println(a.balance)

	fmt.Println("---Ex_23---")
	var sr SalesReport
	sr.AddSale("Burger", 22)
	sr.AddSale("Pizza", 44)
	sr.AddSale("Fries", 4)
	fmt.Println(sr.TotalSales())

	fmt.Println("---Ex_22---")
	a2 := Account2{
		balance: 24444,
		history: []string{},
	}
	a2.Deposit(2344)
	a2.Withdraw(21414)
	fmt.Println(a2.TransactionHistory())

	fmt.Println("---Ex_21---")
	d := Discount{
		persentage: 12,
	}
	o := Order{
		id:       42,
		product:  "Bread",
		quantity: 1000,
		price:    5,
	}

	o.ApplyDiscount(d)
	fmt.Println(o.TotalAmount())

	fmt.Println("---Ex_20---")
	var ml MovieList
	ml.AddMovie("Interstellar", 8.23)
	ml.AddMovie("Avengers: End Game", 7.54)
	ml.AddMovie("The Lord of Rings: Return of the King", 8.89)
	fmt.Println(ml.AverageRating())
}
