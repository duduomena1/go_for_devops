package main

import (
	"fmt"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	john := Person{Name: "John", Age: 25}
	fmt.Println("Declared and instantiated together:", john)

	// using the new keyword

	jane := new(Person)
	jane.Name = "Jane"
	jane.Age = 30
	fmt.Println("Using new keyword:", *jane)

	anon := struct {
		country string
		code    int
	}{
		country: "Brasil",
		code:    55,
	}
	fmt.Println("Anonymous struct:", anon)

	//nested struct
	type Address struct {
		City  string
		State string
	}
	type Employee struct {
		PersonalInfo Person
		Address      Address
	}

	emp := Employee{
		PersonalInfo: Person{Name: "Alisson", Age: 23},
		Address:      Address{City: "Salvador", State: "BA"},
	}
	fmt.Println("Nested struct:", emp)

	//Embedded Anonymous fields
	type Manager struct {
		Person
		Department string
	}
	mgr := Manager{
		Person:     Person{Name: "Eduardo", Age: 27},
		Department: "IT",
	}
	fmt.Println("Embedded Anonymous fields:", mgr)
}
