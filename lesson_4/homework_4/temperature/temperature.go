package temperature

import (
	"fmt"
	"os"
	"strconv"
)

func TemperatureDescription() {
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a degree`)
		return
	}

	temperature := os.Args[1]
	intTemperature, err := strconv.Atoi(temperature)
	if err != nil {
		fmt.Println("Error")
	}

	fmt.Println(getTemperatureDescription(intTemperature))
}

func getTemperatureDescription(temperature int) string {
	if temperature < 10 {
		return "Cold"
	} else if temperature >= 10 && temperature <= 25 {
		return "Warm"
	} else {
		return "Hot"
	}
}
