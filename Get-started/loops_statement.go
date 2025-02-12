package main

import (
	"fmt"
)

func main() {
	fmt.Println("Basic for Loops")
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("for as a while loop")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	//fmt.Println("inifinite for loop")
	//for {
	//	if someCondition {
	//		break
	//	}
	//}
	fmt.Println("for with range") //for slices, arrays, maps, strings
	nums := []int{2, 3, 5, 7, 11, 13}
	for index, value := range nums {
		fmt.Println("index: %d, value: %d\n", index, value)
	}
	fmt.Println("for loop with range") //for strings
	for Index, runeValue := range "Hello" {
		fmt.Printf("Index: %d, Rune: %c\n", Index, runeValue)
	}
}
