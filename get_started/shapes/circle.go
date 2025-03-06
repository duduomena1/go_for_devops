package shapes

import "math"

func CircleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func CircleDiameter(radius float64) float64 {
	return 2 * radius
}
