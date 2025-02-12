package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices Examples")
	fmt.Println("Using slices without initialization:")

	var slice1 []int
	fmt.Println("Using slices with initialization:")
	fmt.Println("Slice without initialization:", slice1)
	slice2 := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice with Initialization:", slice2)

	fmt.Println("Slices using Make Function:")
	slice3 := make([]int, 5)
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 5
	fmt.Println("Slice using Make Function:", slice3)
	slice4 := make([]int, 3, 5)
	slice4[0] = 9
	slice4[1] = 6
	slice4[2] = 5
	fmt.Println("slice 4:", slice4)

	slice4 = append(slice4, 7)
	slice4 = append(slice4, 8)
	fmt.Println("slice 4 after append:", slice4)
	fmt.Println("slice 4 capacity:", cap(slice4))
	slice4 = append(slice4, 10)
	slice4 = append(slice4, 11)
	fmt.Println("slice 4 after append again:", slice4)
	fmt.Println("slice 4 capacity:", cap(slice4), ",length of slice:", len(slice4))

	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice5 := array[2:7] //before the : include and after : exclude
	fmt.Println("slice 5", slice5)

}
