package greeting

import (
	"fmt"
	"os"
)

func Greeting() {
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a day type:
		morning, noon, evening`)
		return
	}

	dayType := os.Args[1]
	printGreeting(dayType)
}

func printGreeting(dayType string) {
	switch dayType {
	default:
		fmt.Println(`Please provide a day type:
		morning, noon, evening`)
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
