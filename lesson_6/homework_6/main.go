package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var amount int = 100000
	fmt.Println("---addUpToAmount()---")
	addUpToAccount(&amount)
	fmt.Println(amount)

	var loanAmount int = 3000000
	fmt.Println("---payLoan()---")
	payLoan(&loanAmount)
	fmt.Println(loanAmount)

	*&amount = 500000
	fmt.Println("---transferMoney()---")
	transferMoney(&amount)
	fmt.Println(amount)

	*&amount = 100000
	fmt.Println("---accumDeposit()---")
	accumDeposit(&amount)
	fmt.Println(amount)

	var currency float64 = 74.5
	fmt.Println("---convertToUSD()---")
	convertToUSD(&currency)
	fmt.Println(int(currency))

	*&amount = 200000
	fmt.Println("---accumCompoudInterest()---")
	accumCompoundInterest(&amount)

	*&amount = 5000
	fmt.Println("---decrementLimitedAmount()---")
	decrementLimitedAmount(&amount)

	*&amount = 500000
	fmt.Println("---accumYearlyDeposit()---")
	accumYearlyDeposit(&amount)

	*&amount = 100000
	fmt.Println("---transactWithComission()---")
	transactWithCommision(&amount)

	*&amount = 300000
	fmt.Println("---profitFromInvestment()---")
	profitFromInvestment(&amount)

}

func addUpToAccount(amount *int) {
	for *amount < 500000 {
		*amount += 10000
	}

	fmt.Println("Limit was reached")

}

func payLoan(amount *int) {
	for *amount > 500000 {
		*amount -= (*amount / 12) + (*amount / 100 * 10)
	}

	fmt.Println("The loan is almost paid")
}

func transferMoney(account *int) {
	var amount int
	fmt.Print("Transfer amount >>> ")
	fmt.Scan(&amount)
	if amount <= 100000 {
		*account += amount
		return
	}

	fmt.Println("Transaction amount is above the limit")
}

func accumDeposit(amount *int) {
	var monthlyAddition = *amount / 100 * 5
	for month := 1; *amount <= 150000; month++ {
		*amount += monthlyAddition
		if *amount > 150000 {
			break
		}

		fmt.Printf("Month %d: %d \n", month, *amount)
	}

	*amount = 150000
	fmt.Println("Deposit limit was reached")
	fmt.Println("Amount after additions:", *amount)
}

func convertToUSD(currency *float64) {
	rand.Seed(time.Now().UnixNano())
	currencyChange := (rand.Float64() * 2) - 1
	if *currency < float64(70) {
		fmt.Println("Currency is too low")
		return
	}

	*currency += currencyChange

	var amount float64
	fmt.Print("Amount of RU to convert >>> ")
	fmt.Scan(&amount)
	fmt.Printf("$%f", amount / *currency)
}

func accumCompoundInterest(amount *int) {
	for month := 1; *amount <= 300000; month++ {
		*amount += *amount / 100 * 5
		if *amount >= 300000 {
			break
		}

		fmt.Printf("Month %d: %d \n", month, *amount)
	}

	*amount = 300000
	fmt.Println("Addition limit was reached")
	fmt.Println("Amount after additions:", *amount)
}

func decrementLimitedAmount(amount *int) {
	var decrementionAmount int
	fmt.Print("Amount to decrement >>> ")
	fmt.Scan(&decrementionAmount)

	if *amount-decrementionAmount < 0 {
		fmt.Println("Daily limit was reached")
		return
	}
	*amount -= decrementionAmount
	fmt.Println(*amount, "remaining")
	decrementLimitedAmount(*&amount)
}

func accumYearlyDeposit(amount *int) {
	var yearlyAddition = *amount / 100 * 6
	for year := 1; *amount <= 700000; year++ {
		*amount += yearlyAddition
		if *amount >= 700000 {
			break
		}

		fmt.Printf("Year %d: %d \n", year, *amount)
	}

	*amount = 700000
	fmt.Println("Amount after additions:", *amount)
}

func transactWithCommision(amount *int) {
	var transactionAmount int
	fmt.Print("Transaction amount >>> ")
	fmt.Scan(&transactionAmount)

	if *amount-transactionAmount+transactionAmount/100 < 50000 {
		fmt.Println("Balance is too low")
		return
	}

	*amount -= transactionAmount + transactionAmount/100
	transactWithCommision(*&amount)
}

func profitFromInvestment(amount *int) {
	var yearlyProfit int = *amount / 100 * 7
	var totalProfit int
	for year := 1; totalProfit <= 400000; year++ {
		totalProfit += yearlyProfit
		if totalProfit >= 400000 {
			fmt.Println("Profit amount was reached")
			break
		}

		fmt.Printf("Year %d: %d \n", year, totalProfit)
	}

	totalProfit = 700000
	fmt.Println("Profit after addtions", totalProfit)
}
