package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 5, 7, 2, 5, 23, 5, 3}
	fmt.Println("Original Slice:", slice)

	fmt.Printf("length: %d, capacity %d\n", len(slice), cap(slice))

	slice = append(slice, 7)
	fmt.Println("Slice after appending", slice)

	subSlice := slice[2:5]
	fmt.Println("Subslice:", subSlice)

	slice = append(slice[:4], slice[5:]...)
	fmt.Println("Slice after removing element", slice)

	if len(slice) == 0 {
		fmt.Println("Slice is empty")
	} else {
		fmt.Println("Slice is not empty")
	}

	sort.Ints(slice)
	fmt.Println("Sorted Slice:", slice)

	elementToFind := 43
	found := false
	for _, v := range slice {
		if v == elementToFind {
			found = true
			break
		}
	}
	if found {
		fmt.Printf("Element %d found in the slice\n", elementToFind)
	} else {
		fmt.Printf("Element %d not Found in the slice\n", elementToFind)
	}
}
