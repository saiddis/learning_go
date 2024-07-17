package trafficlight

import (
	"fmt"
	"os"
)

func TrafficLight() {
	if len(os.Args) < 2 {
		fmt.Println(`Please provide a traffic light color:
		red, yellow, green`)
		return
	}

	lightColor := os.Args[1]
	printTrafficLight(lightColor)
}

func printTrafficLight(lightColor string) {
	switch lightColor {
	default:
		fmt.Println(`Please provide a traffic light color:
		red, yellow, green`)
	case "red":
		fmt.Println("Stop")
	case "yellow":
		fmt.Println("Watch out")
	case "green":
		fmt.Println("Go")
	}
}
