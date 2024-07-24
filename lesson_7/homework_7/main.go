package main

import (
	"fmt"
)

func main() {
	fmt.Println("---assessTemperature()---")
	var degree Temperature = 35
	assessTemperature(degree)

	fmt.Println("---isTeenager(Age)---")
	var age Age = 34
	isTeenager(age)

	fmt.Println("---checkSpeed(Speed)---")
	var speed Speed = 45
	checkSpeed(speed)

	fmt.Println("---checkScore(Score)---")
	var score Score = 213
	checkScore(score)

	fmt.Println("---assessWeightByBMI(BMI)---")
	var massIndex BMI = 23
	assessWeightByBMI(massIndex)

	fmt.Println("---UnaryOperation(int)int---")
	var doubledNum UnaryOperation = doubleNum
	fmt.Println(doubledNum(3))

	fmt.Println("---Check(int)int---")
	var checkEven Check = isEven
	fmt.Println(checkEven(123))
	var checkPositive Check = isPositive
	fmt.Println(checkPositive(12))

	fmt.Println("---Operator(int int)int---")
	var a, b int = 3, 5
	var summedNums, diffedNums, multipliedNums Operator = sum, diff, mul
	fmt.Println(summedNums(a, b))
	fmt.Println(diffedNums(a, b))
	fmt.Println(multipliedNums(a, b))

	fmt.Println("---reverseCounter(Count)---")
	var numCount Count = 3
	reverseCounter(numCount)

	fmt.Println("---displayBankAccountData(BankAccount)---")
	var user AccountType = "user"
	bankAccount := BankAccount{
		"said",
		user,
		4323,
	}
	displayBankAccountData(bankAccount)

	fmt.Println("---displayOrderData(Order)---")
	var status OrderStatus = "pending"
	order := Order{
		323,
		status,
		9,
	}
	displayOrderData(order)

	fmt.Println("---displayVehicleData(Vehicle)---")
	var vehicleType VehicleType = "van"
	vehicle := Vehicle{
		"somebody",
		vehicleType,
		123879,
	}
	displayVehicleData(vehicle)

	fmt.Println("---checkTransactions([]bool) bool)---")
	var trans = []float64{23.43, 323.45, 7745}
	fmt.Println(checkTransactions(trans))

	fmt.Println("---displayBookData(Book)---")
	book := Book{
		"Learning go",
		"Jon Bodner",
		434,
	}
	displayBookData(book)
}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func displayBookData(book Book) {
	fmt.Printf("'%v' has %d pages \n", book.Title, book.Pages)
}

type InventoryID = int

type Transaction = func(float64) bool

func makeTransaction(amount float64) bool {
	var tr Transaction = func(amount float64) bool {
		if amount > float64(0) {
			return true
		}

		return false
	}

	return tr(amount)
}

func checkTransactions(trans []float64) bool {
	for _, v := range trans {
		if makeTransaction(v) {
			return true
		}
	}

	return false
}

type VehicleType = string
type Vehicle struct {
	Owner   string
	Type    VehicleType
	Mileage int
}

func displayVehicleData(vehicle Vehicle) {
	fmt.Println("Owner:", vehicle.Owner)
	fmt.Println("Type:", vehicle.Type)
	fmt.Println("Mileage:", vehicle.Mileage)
}

type OrderStatus = string
type Order struct {
	ID     int
	Status OrderStatus
	Amount int
}

func displayOrderData(order Order) {
	fmt.Println("ID:", order.ID)
	fmt.Println("Status:", order.Status)
	fmt.Println("Amount:", order.Amount)
}

type AccountType = string
type BankAccount struct {
	Owner   string
	Type    AccountType
	Balance int
}

func displayBankAccountData(account BankAccount) {
	fmt.Println("Owner:", account.Owner)
	fmt.Println("Type:", account.Type)
	fmt.Println("Balance:", account.Balance)
}

type Count = int

func reverseCounter(num Count) {
	for i := num; i >= 0; i-- {
		fmt.Println(i)
	}
}

type Operator = func(int, int) int

func sum(a, b int) int {
	return a + b
}

func diff(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

type Check = func(int) bool

func isPositive(num int) bool {
	if num >= 0 {
		return true
	} else {
		return false
	}
}

func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}

	return false
}

type UnaryOperation = func(int) int

func doubleNum(num int) int {
	return num * 2
}

type BMI = float64

func assessWeightByBMI(index BMI) {
	if index < 18.5 {
		fmt.Println("Under weight")
	} else if index >= 18.5 && index < 25 {
		fmt.Println("Normal weight")
	} else if index >= 25 && index < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obesity")
	}
}

type Score = int

func checkScore(score Score) {
	if score > 0 {
		fmt.Println("The score is negative")
	} else if score > 0 {
		fmt.Println("The score is positive")
	} else {
		fmt.Println("The score is zero")
	}
}

type Speed = float64

func checkSpeed(speed Speed) {
	if speed >= 0 && speed <= 120 {
		fmt.Println("Keep going")
	} else if speed > 120 {
		fmt.Println("Slow down!")
	}
}

type Age = int

func isTeenager(age Age) {
	if age >= 13 && age <= 19 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Not a teenager")
	}
}

type Temperature = float64

func assessTemperature(num Temperature) {
	if num > 0 {
		fmt.Println("Above 0")
	} else if num < 0 {
		fmt.Println("Under 0")
	} else {
		fmt.Println("0")
	}
}
