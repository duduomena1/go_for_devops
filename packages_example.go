package main

import (
	"fmt"
	"go_for_devops/shapes"
)

func main() {
	radius := 5.0
	areaCircle := shapes.CircleArea(radius)
	fmt.Printf("Area of circle: %f\n", areaCircle)

	side := 3.0
	areasquare := shapes.SquareArea(side)
	fmt.Printf("Area of square: %f\n", areasquare)

	diameter := shapes.CircleDiameter(radius)
	fmt.Printf("Diameter of circle: %f\n", diameter)
}
