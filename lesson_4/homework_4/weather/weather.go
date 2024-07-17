package weather

import (
	"fmt"
	"os"
)

func Weather() {
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a weather type:
		sunny, clody, rainy`)
		return
	}

	weatherType := os.Args[1]
	printWeather(weatherType)
}

func printWeather(weatherType string) {
	switch weatherType {
	default:
		fmt.Println(`Please provide a weather type:
		sunny, clody, rainy`)
	case "sunny":
		fmt.Println("Sunny")
	case "cloudy":
		fmt.Println("Cloudy")
	case "rainy":
		fmt.Println("Rainy")
	}
}
