package main

import (
	"fmt"
)

func modifyArray(arr [5]int, index int, value int) [5]int {
	arr[index] = value
	return arr
}

func main() {
	originalArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Length of arrays is:", len(originalArray))

	copiedArray := originalArray
	fmt.Println("Copied Array:", copiedArray)

	modifiedArray := modifyArray(originalArray, 2, 12)
	fmt.Println("Modified Array:", modifiedArray)

	isEqual := originalArray != modifiedArray
	fmt.Println("Are original and modified arrays different?:", isEqual)

	isCopiedEqual := originalArray == copiedArray
	fmt.Println("Are original and copied arrays equal?", isCopiedEqual)

	fmt.Println("Array iterations")

	numbers := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Iteration using traditional for loop:")
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, numbers[i])
	}
	fmt.Println("\nIteration using range0 loop:")
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
