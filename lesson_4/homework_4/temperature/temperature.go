package temperature

import (
	"fmt"
)

func TemperatureDescription() {
	var temperature int
	fmt.Print(`Type the temperature 
>>> `)
	fmt.Scan(&temperature)
	fmt.Println(getTemperatureDescription(temperature))
}

func getTemperatureDescription(temperature int) string {
	if temperature < 10 {
		return "Cold"
	} else if temperature >= 10 && temperature <= 25 {
		return "Warm"
	} else if temperature > 25 {
		return "Hot"
	} else {
		return "Error: given degree or input type is not valid"
	}
}
