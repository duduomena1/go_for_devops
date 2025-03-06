package main

import (
	"fmt"
)

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Println("Inside modifySlice: after modification", s)
}

func main() {
	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original Slice before modification:", originalSlice)

	modifySlice(originalSlice)

	fmt.Println("Original Slice after modification:", originalSlice)
}
