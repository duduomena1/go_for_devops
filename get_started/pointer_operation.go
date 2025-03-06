package main

import (
	"fmt"
)

func modifyValue(ptr *int) {
	*ptr = 43
}

func main() {
	a := 22
	ptrA := &a
	fmt.Println("Original value of a:", a)
	fmt.Println("Dereferencing ptrA gives:", *ptrA)

	//Changing the value of pointer

	*ptrA = 30
	fmt.Println("After changing the value of pointer, a:", a)

	//Comparing Pointers

	b := 20
	ptrB := &b
	fmt.Println("ptrA and ptrB pont to the same address:", ptrA == ptrB)
	p1 := &a
	p2 := &b
	p3 := &a

	fmt.Println("p1 and p2 point to the same address?", p1 == p2)
	fmt.Println("p1 and p2 to the same address?:", p1 == p3)

	fmt.Println("Value of a before passing to function:", a)
	modifyValue(ptrA)
	fmt.Println("Value of a after passing to function:", a)
}
