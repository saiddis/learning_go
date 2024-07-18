package main

import (
	"discount"
	"grade"
	"greeting"
	"temperature"
	"trafficlight"
	"weather"
)

func main() {
	greeting.Greeting()
	discount.Discount()
	trafficlight.TrafficLight()
	temperature.TemperatureDescription()
	grade.Grade()
	weather.Weather()
}
