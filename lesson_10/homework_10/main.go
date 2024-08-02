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

type Movie struct {
	title  string
	rating float64
}

type MovieList struct {
	Movies []Movie
}

func (ml *MovieList) AddMovie(title string, rating float64) {
	movie := Movie{
		title:  title,
		rating: rating,
	}

	ml.Movies = append(ml.Movies, movie)
}

func (ml MovieList) AverageRating() float64 {
	var totalRating float64
	for _, v := range ml.Movies {
		totalRating += v.rating
	}
	return totalRating / float64(len(ml.Movies))
}

type Discount struct {
	persentage float64
}

func (o *Order) ApplyDiscount(discount Discount) {
	o.price = o.price / float64(100) * discount.persentage
}

func (o Order) TotalAmount() float64 {
	return o.price * float64(o.quantity)
}

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

type Task struct {
	title       string
	description string
	status      string
}

type Project struct {
	Tasks []Task
}

func (p *Project) AddTask(title, description, status string) {
	task := Task{
		title:       title,
		description: description,
		status:      status,
	}

	p.Tasks = append(p.Tasks, task)
}

func (p Project) CountTasksByStatus(status string) int {
	var taskAmount int

	for _, v := range p.Tasks {
		if status == v.status {
			taskAmount++
		}
	}

	return taskAmount
}

type User struct {
	username string
	email    string
	age      int
}

type UserManager struct {
	Users []User
}

func (um *UserManager) AddUser(username, email string, age int) {
	user := User{
		username: username,
		email:    email,
		age:      age,
	}

	um.Users = append(um.Users, user)
}

func (um UserManager) UsersOlderThan(age int) []User {
	var olderUsers []User

	for _, v := range um.Users {
		if v.age > age {
			olderUsers = append(olderUsers, v)
		}
	}

	return olderUsers
}

type Contract struct {
	constractID int
	client      string
	amount      float64
}

type ContractManager struct {
	Contracts []Contract
}

func (cm *ContractManager) AddContract(cID int, client string, amount float64) {
	contract := Contract{
		constractID: cID,
		client:      client,
		amount:      amount,
	}

	cm.Contracts = append(cm.Contracts, contract)
}

func (cm ContractManager) TotalAmountForClient(client string) float64 {
	var totalAmount float64

	for _, v := range cm.Contracts {
		if v.client == client {
			totalAmount += v.amount
		}
	}

	return totalAmount
}

type Book struct {
	title  string
	author string
	year   int
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(title, author string, year int) {
	book := Book{
		title:  title,
		author: author,
		year:   year,
	}

	l.Books = append(l.Books, book)
}

func (l Library) BooksByAuthorAfterYear(year int) []Book {
	var booksAfterYear []Book

	for _, v := range l.Books {
		if v.year > year {
			booksAfterYear = append(booksAfterYear, v)
		}
	}

	return booksAfterYear
}

type Activity struct {
	activityType string
	timeStamp    time.Time
}

type UserActivityTracker struct {
	Activities []Activity
}

func (uat *UserActivityTracker) AddActivity(activityType string, timestamp time.Time) {
	activity := Activity{
		activityType: activityType,
		timeStamp:    timestamp,
	}

	uat.Activities = append(uat.Activities, activity)
}

func (uat UserActivityTracker) ActivityAfterTime(timestamp time.Time) []Activity {
	var actsAfterTime []Activity

	for _, v := range uat.Activities {
		if v.timeStamp.Compare(timestamp) == 1 {
			actsAfterTime = append(actsAfterTime, v)
		}
	}

	return actsAfterTime
}
