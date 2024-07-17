package main

import (
	"fmt"
)

func main() {
	/* num := 36
	if num > 0 && num < 10 {
		fmt.Println("One digital")
	} else if num > 9 && num < 100 {
		fmt.Println("Two digital")
	} else if num > 99 && num < 1000 {
		fmt.Println("Three digital")
	} else if num > 999 && num < 10000 {
		fmt.Println("Four digital")
	}

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	} */

	/* a := 45
	b := 54
	c := 453 */

	/* if a > b {
		fmt.Println(a)
	} else if a < b {
		fmt.Println(b)
	} else {
		fmt.Println("Equal")
	} */

	//if a > b && b > c {
	/* fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c) */
	//fmt.Println(a + b)
	//} else if c > b && b > a {
	/* fmt.Println(c)
	fmt.Println(b)
	fmt.Println(a) */
	//fmt.Println(c + b)
	//} else if b > c && c > a {
	/* fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a) */
	//fmt.Println(b + c)
	/*} else if c > a && a > b {
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(b) */

	/*} else if a > c && c > b {
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(b) */

	/* else {
		fmt.Println(b)
		fmt.Println(a)
		fmt.Println(c)

	}  */

	/* switch {

	case a > b:
		fmt.Println(a)
		fallthrough

	case b > a:
		fmt.Println(b)
		fallthrough

	default:
		fmt.Println(c)
	} */

	/* year := 2024

	if year%4 == 0 || year%4 != 0 && year%400 == 0 {
		fmt.Println(year, true)
	} else {
		fmt.Println(year, false)
	} */

	defineMonth(4)

}

var months = map[int]int{
	1:  31,
	2:  28,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  30,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func defineMonth(monthNum int) {
	fmt.Printf("Month %v has %v days", monthNum, months[monthNum])
}
