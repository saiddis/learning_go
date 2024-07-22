package main

import "fmt"

func main() {
	// buildPascalPyramid()
	/* var x int = 10
	var ptr *int

	ptr = &x */

	/* fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(ptr) */
	// fmt.Println(*ptr)

	var n int = 10
	var m int = 50

	fmt.Println("---update()---")
	fmt.Println(n)
	update(&n)
	fmt.Println(n)

	fmt.Println("---swap()---")
	swap(&n, &m)
	fmt.Println(n, m)

	fmt.Println("---printValue()---")
	printValue(&n)

	fmt.Println("---increment()---")
	increment(&m)
	fmt.Println(m)

	var name string = "Said"
	var age int = 17

	fmt.Println("---printPersonalData()---")
	printPersonalData(&name, &age)

	fmt.Println("---moneyConvertion()---")
	var amont float32 = 3544
	moneyConversion(&amont)
}

func update(num *int) {
	*num = *num + 10
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func printValue(num *int) {

	fmt.Println(*num)
}

func increment(num *int) {
	*num++
}

func printPersonalData(name *string, age *int) {
	fmt.Println("Name:", *name)
	fmt.Println("Age:", *age)
}

func moneyConversion(amount *float32) {
	*amount *= 0.85
	fmt.Println("Conversion from USD to EURO:", *amount)
}

/* func addGrade(subject *string, grade *int) {

}

func printAverageGrade(score *int) int {

} */

/* func buildPascalPyramid() {
	var num int
	fmt.Print("Type a number")
	fmt.Scan(&num)

	fmt.Println(binomialCoefficient(num, 5))
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func binomialCoefficient(n, k int) int {
	return factorial(n) / (factorial(k) * factorial(n-k))
} */
