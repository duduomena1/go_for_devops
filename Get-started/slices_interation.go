package main

import (
	"fmt"
)

func main() {
	slice := []string{"banana", "apple", "orange", "grape", "kiwi"}
	fmt.Println("Original Slice:", slice)
	fmt.Println("iteration with for loop:")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index %d, Value %s\n", i, slice[i])
	}
	fmt.Println("\nIterating with range:")
	for index, value := range slice {
		fmt.Printf("Index %d, Value %s\n", index, value)
	}
	findInSlice := "banana"
	found := false
	for _, v := range slice {
		if v == findInSlice {
			found = true
			break
		}
	}
	if found {
		fmt.Printf("Element %s found in the slice\n", findInSlice)
	} else {
		fmt.Printf("Element %s not found in the slice\n", findInSlice)
	}

}
