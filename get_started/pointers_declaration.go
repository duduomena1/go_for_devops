package main

import (
	"fmt"
)

func main() {
	//using var
	var ptr1 *int
	fmt.Println("1. using var to declare a pointer", ptr1)

	var value int = 43
	ptr2 := &value
	fmt.Println("2. Pointer to en existing Variable: ", *ptr2)

	ptr3 := new(int)
	*ptr3 = 400
	fmt.Println("3. Create a pointer with new: ", *ptr3)

	//pointer to Pointer

	ptr4 := &ptr3
	fmt.Println("4. Pointer to Pointer: ", **ptr4)

	fmt.Println(&ptr3)
	fmt.Println(ptr3)
}
