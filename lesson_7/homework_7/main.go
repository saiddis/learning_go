package main

import (
	"fmt"
	"math/rand"
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
	var tr1, tr2, tr3 Transaction = makeTransaction, makeTransaction, makeTransaction
	var trans = []Transaction{tr1, tr2, tr3}
	fmt.Println(checkTransactions(trans))

	fmt.Println("---displayItemData(Item)---")
	var inventoryID InventoryID = 2345
	var item = Item{
		inventoryID,
		"Laptop",
		"High",
	}
	displayItemData(item)

	fmt.Println("---displayBookData(Book)---")
	book := Book{
		"Learning go",
		"Jon Bodner",
		434,
	}
	displayBookData(book)

	fmt.Println("---updateProductData(Product)---")
	var product = Product{
		"Apple",
		56,
	}
	updateProductData(&product)

	fmt.Println("---printStudentNames(Student)---")
	student1 := &Student{Name: "Umar"}
	student2 := &Student{Name: "Akbar"}
	student3 := &Student{Name: "Said"}
	student1.Next = student2
	student2.Next = student3
	printStudentNames(*student1)

	fmt.Println("---Event---")
	fmt.Println("Event name:", birthday.EventName)
	fmt.Println("Location:", birthday.Location)

	fmt.Println("---getStudent2Data(Student2) (int, string, []float64)---")
	st1 := Student2{
		34234,
		"Said",
		[]float64{4, 5, 3, 5, 4},
	}
	st2 := Student2{
		4543,
		"Umar",
		[]float64{4, 2, 5, 3, 3},
	}
	st3 := Student2{
		3324,
		"Muin",
		[]float64{5, 5, 5, 5, 4},
	}
	printStudent2Data(st1)
	printStudent2Data(st2)
	printStudent2Data(st3)
	students := []Student2{st1, st2, st3}
	fmt.Println("---getTheBestStudent([]Student2) Student2---")
	fmt.Println(getTheBestStudent(students))
}

type Student2 struct {
	ID     int
	Name   string
	Grades []float64
}

func printStudent2Data(st Student2) {
	fmt.Println("ID:", st.ID)
	fmt.Println("Name:", st.Name)
	fmt.Println("Grades:", st.Grades)
}

func getTheBestStudent(students []Student2) Student2 {
	var sortedStudentsIndexesByGrades []int
	for i := len(students) - 1; i > 0; i-- {
		sortedStudentsIndexesByGrades =
			append(sortedStudentsIndexesByGrades, compareStudentsGrades(students[i], students[i-1])+i)
	}

	return students[sortedStudentsIndexesByGrades[0]]
}

func compareStudentsGrades(st1, st2 Student2) int {
	st1AG := getAverageGrade(st1.Grades)
	st2AG := getAverageGrade(st2.Grades)

	if st1AG > st2AG {
		return 0
	}

	return -1
}

func getAverageGrade(grades []float64) float64 {
	var sum float64
	for _, v := range grades {
		sum += v
	}

	return sum / float64(len(grades))
}

var birthday = struct {
	EventName string
	Location  string
}{
	EventName: "birthday",
	Location:  "home",
}

type Student struct {
	Name string
	Next *Student
}

func printStudentNames(student Student) {
	fmt.Println(student.Name)
	if student.Next != nil {
		printStudentNames(*student.Next)
	}
}

type Product struct {
	Name  string
	Stock int
}

func updateProductData(p *Product) {
	if p.Stock < 100 {
		*&p.Stock += 200
		fmt.Printf("Updated %v amount: %d \n", p.Name, p.Stock)
		return
	}

	fmt.Printf("The %v amount is enough \n", p.Name)

}

type Book struct {
	Title  string
	Author string
	Pages  int
}

func displayBookData(book Book) {
	fmt.Printf("'%v' has %d pages \n", book.Title, book.Pages)
}

type InventoryID int
type Item struct {
	ID      InventoryID
	Name    string
	Quality string
}

func displayItemData(item Item) {
	fmt.Println("ID:", item.ID)
	fmt.Println("Name:", item.Name)
	fmt.Println("Quality:", item.Quality)
}

type Transaction func(float64, float64) bool

func makeTransaction(balance, price float64) bool {
	if balance > price {
		return true
	}
	return false
}

func checkTransactions(trans []Transaction) bool {
	var balance, price float64
	for _, v := range trans {
		balance = rand.Float64() * float64(rand.Intn(1000))
		price = rand.Float64() * float64(rand.Intn(100))
		if v(balance, price) {
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
