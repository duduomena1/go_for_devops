package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
)

func main() {
	// Basic Arithmetic Operations

	var a, b int = 10, 5
	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Modulus:", a%b)

	//working with float numbers

	var c, d float64 = 10.5, 5.5
	fmt.Println("Float Addition:", c+d)
	fmt.Println("Float Subtraction:", c-d)
	fmt.Println("Float Multiplication:", c*d)
	fmt.Println("Float Division:", c/d)

	//Type conversion
	var e int = 42
	var f float64 = float64(e)
	fmt.Println("Int to float:", f)

	var g float64 = 46.59856
	var h int = int(g)
	fmt.Println("Float to int:", h)

	//Complex numbers

	var complex1 complex128 = complex(12, 22)
	var complex2 complex128 = complex(5, 7)
	fmt.Println("Complex Addition:", complex1+complex2)
	fmt.Println("Complex Subtraction:", complex1-complex2)
	fmt.Println("Complex Multiplication:", complex1*complex2)
	fmt.Println("Complex Division:", complex1/complex2)
	fmt.Println("Complex Absolute Value:", cmplx.Abs(complex1))

	// Math functions

	fmt.Println("Power:", math.Pow(2, 3))
	fmt.Println("Square Root:", math.Sqrt(16))
	fmt.Println("sin", math.Sin(math.Pi/2))
	fmt.Println("cos:", math.Cos(math.Pi))
	fmt.Println("Log:", math.Log(math.E))

	// String to number conversion

	numStr := "3.1415"
	parsedFloat, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		fmt.Println("Error while parsing float:", err)
	} else {
		fmt.Println("Parsed Float:", parsedFloat)
	}
}
