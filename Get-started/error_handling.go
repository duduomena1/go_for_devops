package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	msg string
}

func (e *CustomError) Error() string {
	return e.msg
}

func SomeFunction(flag bool) error {
	if !flag {
		return &CustomError{"Custom error occurred"}
	}
	return nil
}
func SafeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	panic("A problem occurred")
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}
func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	err2 := SomeFunction(false)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
	SafeFunction()
}
