package main

import (
	"fmt"
)

type person struct {
	Name     string
	Age      int
	Salary   float64
	Position string
}

func updateSalary(emp *person, newsalary float64) {
	emp.Salary = newsalary
}

func main() {
	emp := person{
		Name:     "Eduardo",
		Age:      25,
		Salary:   10000,
		Position: "DevOps",
	}
	fmt.Println("Salary=", emp.Salary)
	updateSalary(&emp, 15000)

	fmt.Println("After the promotion", emp.Name, "starts receive:", emp.Salary)
}
