package main

import (
	"fmt"
)

func main() {
	var ar1 [5]int
	fmt.Println("Array  declared without initialization:", ar1)
	var ar2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array  declared with initialization:", ar2)
	ar3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array with sort hand declaration:", ar3)
	ar4 := [5]int{1: 10, 3: 30, 0: 2}
	fmt.Println("Array initializing with specific elements:", ar4)
	ar5 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Array using ... operator dor infer length:", ar5)

	var multArray [3][4]int
	multArray[0] = [4]int{1, 2, 3}
	multArray[1] = [4]int{4, 5, 6}
	fmt.Println("Multi-dimensional Array:", multArray)

	initializedMultiArray := [2][3]int{{7, 8, 9}, {10, 11, 12}}
	fmt.Println("Initialized Multi-dimensional Array:", initializedMultiArray)

	var pointerArray [3]*int
	num1, num2, num3 := 13, 14, 15
	pointerArray[0] = &num1
	pointerArray[1] = &num2
	pointerArray[2] = &num3
	fmt.Println("Array with pointers:")
	for i := 0; i < len(pointerArray); i++ {
		fmt.Printf("pointerArray[%d] = %d\n", i, *pointerArray[i])
	}

	type Person struct {
		Name string
		Age  int
	}
	var people [2]Person
	people[0] = Person{"Eduardo", 25}
	people[1] = Person{"Tiaguinho", 50}
	fmt.Println("Array of Structs:", people)
	fmt.Println(people[0:1])
}
