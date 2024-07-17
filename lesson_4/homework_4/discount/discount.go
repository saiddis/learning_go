package discount

import (
	"fmt"
	"os"
	"strconv"
)

func Discount() {
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a price`)
		return
	}

	amount := os.Args[1]
	intAmount, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(getDiscount(intAmount))
}

func getDiscount(amount int) string {
	if amount > 1000 {
		return "10%"
	} else if amount >= 100 && amount < 500 {
		return "5%"
	} else {
		return "0%"
	}
}
