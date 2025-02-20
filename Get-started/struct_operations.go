package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Printf("Hello, my name is %s and i am %d years old\n", p.Name, p.Age)
}

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	alice := Person{Name: "Icaro", Age: 21}
	fmt.Println("Struct Instance:", alice)

	//accessing and modify fields

	fmt.Println("Name:", alice.Name)
	alice.Name = "Juliana Santos"
	fmt.Println("Name after 6:00PM:", alice.Name)

	bob := &Person{Name: "Bob", Age: 30}
	fmt.Println("Pointer to a struct:", *bob)

	alice.SayHello()
	bob.SayHello()

	emp := &Employee{ID: 1, Name: "John Doe"}
	t := reflect.TypeOf(*emp)
	field, _ := t.FieldByName("Name")
	fmt.Println("Tag on Name field:", field.Tag.Get("json"))

	type Address struct {
		City  string
		State string
	}
	city := &Address{City: "Salvador", State: "BA"}
	fmt.Println("Address struct:", city)

	type job struct {
		Role     string   `json:"Role"`
		Employee Employee `json:"employee"`
		person   Person   `json:"person"`
		address  Address  `json:"Address"`
	}
	role := job{"devops", *emp, alice, *city}

	fmt.Println("Job struct:", role)
}
