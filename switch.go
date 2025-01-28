package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Using basic Switch case:")
	day := 3
	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	case 5:
		fmt.Println("friday")
	case 6:
		fmt.Println("saturday")
	case 7:
		fmt.Println("sunday")
	default:
		fmt.Println("invalid day")
	}
	fmt.Println("Switch with Initialization")
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}
	fmt.Println("using switch without a Condition:")
	x := 42
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	default:
		fmt.Println("x is positive")
	}
}
