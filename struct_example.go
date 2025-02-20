package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SensorData struct {
	Temperature float64
	Humidity    float64
	Pressure    float64
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var sensorRead [3]SensorData

	for i := range sensorRead {
		sensorRead[i] = SensorData{
			Temperature: rand.Float64()*25 + 15,
			Humidity:    rand.Float64() * 100,
			Pressure:    rand.Float64()*20 + 980,
		}
	}
	for i, reading := range sensorRead {
		fmt.Printf("Sensor %d - Temperature: %.2f°C, Humidity: %.2f%%, Pressure: %.2f hPa\n",
			i+1, reading.Temperature, reading.Humidity, reading.Pressure)
	}
}
