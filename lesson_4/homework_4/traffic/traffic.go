package trafficlight

import (
	"fmt"
)

func TrafficLight() {
	var lightColor string
	fmt.Print(`Type the traffic light color: red, yellow, green 
>>> `)
	fmt.Scan(&lightColor)
	printTrafficLight(lightColor)
}

func printTrafficLight(lightColor string) {
	switch lightColor {
	default:
		fmt.Println(`Error: given traffic light or input type is not valid`)
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Watch out")
	case "green":
		fmt.Println("Go")
	}
}
