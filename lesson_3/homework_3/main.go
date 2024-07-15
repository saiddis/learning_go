package main

import (
	"fmt"
)

func main() {
	PrintGreeting()
	DisplaySeparator()
	NotifyStart()

	fmt.Println(getWelcomeMessage())
	fmt.Println(getPiValue())
	fmt.Println(IsServerActive())

	PrintNumber(7)
	GreetUser("Said")
	ToggleLight(true)

	adder := func(a, b int) int {
		return a + b
	}
	fmt.Println(adder(1, 2))

	concatenator := func(a, b string) string {
		return a + b
	}
	fmt.Println(concatenator("go", "lang"))

	isPositive := func(num int) bool {
		return num > 0
	}
	fmt.Println(isPositive(3))

	fmt.Println(Calculate(2, 8, adder))
	Execute(false, ToggleLight)
	fmt.Println(Apply(4, func(i int) int { return i }))

	mul := Multiplier(9)
	greetings := StringRepeater(3)
	light_switcher := ConditionalPrinter(true)
	fmt.Println(mul())
	greetings()
	light_switcher()
}

func PrintGreeting() {
	fmt.Println("Hello, World!")
}

func DisplaySeparator() {
	fmt.Println("-------------------")
}

func NotifyStart() {
	fmt.Println("Program startet")
}

func getWelcomeMessage() string {
	return "Welcome!"
}

func getPiValue() float32 {
	return 3.14
}

func IsServerActive() bool {
	return true
}

func PrintNumber(num int) {
	fmt.Println(num)
}

func GreetUser(name string) {
	fmt.Println("Hi, " + name)
}

func ToggleLight(light bool) {
	if light {
		fmt.Println("Light on")
	} else {
		fmt.Println("Light off")
	}
}

func Calculate(a, b int, f func(a, b int) int) int {
	return f(a, b)
}

func Execute(boolean bool, f func(boolean bool)) {
	f(boolean)
}

func Apply(num int, f func(int) int) int {
	return f(num)
}

func Multiplier(num int) func() int {
	return func() int {
		return num * num
	}
}

func StringRepeater(n int) func() {
	greeting := "Great to see you!"
	return func() {
		for i := 0; i < n; i++ {
			fmt.Println(greeting)
		}
	}
}

func ConditionalPrinter(boolean bool) func() {
	return func() {
		ToggleLight(boolean)
	}
}
