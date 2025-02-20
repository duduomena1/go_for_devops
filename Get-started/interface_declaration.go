package main

import (
	"fmt"
)

// Single-method interface
type speaker interface {
	speak() string
}

type Dog struct {
	name string
}

func (d Dog) speak() string {
	return d.name + " says Woof"
}

// Embedding interfaces

type Animal interface {
	speaker
	move() string
}

type Cat struct {
	name string
}

func (c Cat) speak() string {
	return c.name + " says Meow"
}

func (c Cat) move() string {
	return c.name + " Moves quietly"
}

// Empty interface

func PrintAll(v interface{}) {
	fmt.Println(v)
}

type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Start() string {
	return c.Model + " has started"
}

func (c Car) Stop() string {
	return c.Model + " has stopped"
}

func Describe[T speaker](t T) {
	fmt.Println(t.speak())
}

func main() {
	dog := Dog{name: "Godofredo"}
	cat := Cat{name: "Tapioca"}
	car := Car{Model: "UP! TSI"}

	fmt.Println(dog.speak())
	fmt.Println(cat.speak())
	fmt.Println(cat.move())

	PrintAll("An empty interface can hold any data type")
	PrintAll(42)

	fmt.Println(car.Start())
	fmt.Println(car.Stop())

	Describe(dog)
	Describe(cat)

}
