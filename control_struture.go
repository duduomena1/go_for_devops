package main

import (
	"fmt"
)

func main() {
	fmt.Println("using basic If")
	x := 10
	if x > 5 {
		fmt.Println("x is greater than 5")
	}
	fmt.Println("using if-else")

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else {
		fmt.Println("x is less than 5")
	}

	fmt.Println("use if-else else-if")
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x == 10 {
		fmt.Println("x is equal to 10")
	} else {
		fmt.Println("x is less than 10")
	}
	fmt.Println("Using with Initialization Statement")
	if y := 10; y > 5 {
		fmt.Println("y is greater than 5")
	} else {
		fmt.Println("y is less than 5")
	}
}
