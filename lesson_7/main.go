package main

import (
	"fmt"
)

func main() {
	type Person struct {
		name string
		age  int
	}

	/* 	var said = Person{"Said", 17}
	 */
	/* fmt.Println(said)

	said = Person{age: 17, name: "Said"}
	fmt.Println(said)

	fmt.Println(said.name) */

	/* type Employee struct {
		firstName, lastName string
		age, salary         int
	}

	emp := &Employee{"Said", "Islomov", 17, 0}

	fmt.Println((*emp).firstName) */

	/* fmt.Println(isAduld(said.age))
	fmt.Println(isEven(said.age))
	fmt.Println(isInRange(said.age))
	*/
	/* 	var num1, num2 int
	 */
	/* fmt.Print("Type first number >>> \n")
	fmt.Scan(&num1)
	fmt.Print("Type second number >>> \n")
	fmt.Scan(&num2)
	var op = Operation{
		func sum(a, b int) int {
			return a + b
		},

		func diff(a, b int) int {
			return a - b
		},

		func mul(a, b int) int {
			return a * b
		}
	} */

	/* var num Count = 10
	reverseCounter(num)

	var num1 int = 34
	var num2 int = 23
	var checkNums Comparator = isBigger
	fmt.Println(checkNums(num1, num2))

	var mulNums UnaryOperation = doubleNums
	fmt.Println(mulNums(num)) */

	var degree Temperature = 34
	assessTemperature(degree)

	var said = User{"said", 17}

	displayUserData(said)
}

type User struct {
	name string
	age  int
}

func displayUserData(user User) {
	fmt.Println("Name:", user.name)
	fmt.Println("Age:", user.age)
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

type UnaryOperation func(int) int

func doubleNums(num int) int {
	return num * 2
}

func tripleNums(num int) int {
	return num * 3
}

type Comparator func(int, int) bool

func isBigger(a, b int) bool {
	if a > b {
		return true
	}

	return false
}

func isLower(a, b int) bool {
	if a < b {
		return true
	}

	return false
}

type Count = int

func reverseCounter(num Count) {
	for i := num; i >= 0; i-- {
		fmt.Println(i)
	}
}

func isAduld(age int) string {
	if age >= 18 {
		return "The person is adult"
	}

	return "The person is not adult"
}

func isEven(num int) string {
	if num%2 == 0 {
		return "The number is even"
	}

	return "The number is odd"
}

func isInRange(num int) string {
	if num >= 0 && num <= 100 {
		return "The number is in range"
	}

	return "The number is not in range"
}

/* func Operation(a, b *int) {
	fmt.Println("Sum:", sum(*a, *b))
	fmt.Println("Diffrence:", diff(*a, *b))
	fmt.Println("Multiplication:", mul(*a, *b))
} */

type Operation func(a, b int) int
