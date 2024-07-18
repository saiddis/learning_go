package weather

import (
	"fmt"
)

func Weather() {
	var weatherType string
	fmt.Print(`Type the weather: sunny, clody, rainy 
>>> `)
	fmt.Scan(&weatherType)
	printWeather(weatherType)
}

func printWeather(weatherType string) {
	switch weatherType {
	default:
		fmt.Println(`Error: given weather or input type is not valid`)
	case "sunny":
		fmt.Println("Sunny")
	case "cloudy":
		fmt.Println("Cloudy")
	case "rainy":
		fmt.Println("Rainy")
	}
}
