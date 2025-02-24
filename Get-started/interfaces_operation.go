package main

import (
	"fmt"
)

type Greeter interface {
	Greet() string
}

type person struct {
	Name string
}

func (p person) Greet() string {
	return "Hello, " + p.Name
}

func PrintDetail(i interface{}) {
	str, ok := i.(string)
	if ok {
		fmt.Println("It's a string:", str)
	} else {
		fmt.Println("Not a string")
	}
}
func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	people := person{Name: "Eduardo"}
	var greeter Greeter = people
	fmt.Println(greeter.Greet())

	PrintDetail("HEY")
	PrintDetail(10)

	Describe("Hello")
	Describe(2025)
	Describe(3.14)
}
