package main

import (
	"bufio"
	"fmt"
)

func main() {
	fmt.Println("Composite Types")
	fmt.Println("Arrays")
	var o [5]int = [5]int{1, 2, 3, 4}
	fmt.Println("o=", o)

	fmt.Println("Slices") //dynamic-sizes
	p := []int{1, 2, 3, 4, 5}
	fmt.Println("p=", p)

	fmt.Println("structs")
	type person struct {
		Name string
		Age  int
	}
	var p1 person = person{"Dudu", 25}
	fmt.Println("p1 Name=", p1.Name)
	fmt.Println("p1 Age=", p1.Age)
	fmt.Println("p1=", p1)

	fmt.Println("Maps") // for store data in key-value

	m := map[string]int{"Dudu": 25, "Eduardo": 26}
	fmt.Println("m=", m)
	fmt.Println("Dudu=", m["Dudu"])
	fmt.Println("Eduardo=", m["Eduardo"])

	fmt.Println("Pointers") //holds the memory address variable

	s := 10
	t := &s

	fmt.Println("s=", s)
	fmt.Println("t=", t)
	fmt.Println("*t=", *t)

	fmt.Println("Channels") //communication between goroutine, They can be sync or async

	u := make(chan int)
	fmt.Println("u=", u)

	fmt.Println("Functions as Variables")

	var v func(int) int
	v = func(x int) int { return x * x }
	result := v(5)
	fmt.Println("result=", result)

}
