package greeting

import (
	"fmt"
)

func Greeting() {
	var dayType string
	fmt.Print(`Type the daytime: morning, noon, evening
>>> `)
	fmt.Scan(&dayType)
	printGreeting(dayType)
}

func printGreeting(dayType string) {
	switch dayType {
	default:
		fmt.Println(`Error: given daytime or input type is not valid`)
	case "morning":
		fmt.Println("Good morning!")
	case "noon":
		fallthrough
	case "afternoon":
		fmt.Println("Good afternoon!")
	case "evening":
		fmt.Println("Good evening!")
	case "night":
		fmt.Println("Good night!")
	}
}
