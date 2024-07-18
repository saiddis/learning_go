package discount

import (
	"fmt"
)

func Discount() {
	var amount int
	fmt.Print(`Type an amount of money
>>> `)
	fmt.Scan(&amount)
	fmt.Println("Your discount:", getDiscount(amount))
}

func getDiscount(amount int) string {
	if amount >= 1000 {
		return "10%"
	} else if amount >= 500 && amount < 1000 {
		return "5%"
	} else if amount >= 0 && amount > 500 {
		return "0%"
	} else {
		return "Error: given amount or input type is not valid"
	}
}
