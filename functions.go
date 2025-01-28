package main

import (
	"fmt"
)

func add(x, y int) int { //basics function
	return x + y
}
func split(sum int) (x, y int) { // Returns values can be named and used as var defined as te top os function
	x = sum * 4 / 9
	y = sum - x
	return
}
func swap(x, y string) (string, string) { //go can returns multiple values
	return y, x
}

func sum_of_num(nums ...int) int { //can be called with any number os trailing arguments
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	fmt.Println("Basic functions")
	result := add(5, 4)
	fmt.Println("Result=", result)

	fmt.Println("Functions with multiple return values")
	x, y := swap("eduardo", "omena")
	fmt.Println("y, x=", x, y)

	fmt.Println("Named return values:")
	sum := 100
	k, l := split(sum)
	fmt.Printf("The split of %d is: k = %d, l = %d\n", sum, k, l)

	fmt.Println("variadic functions:")
	sum_result := sum_of_num(1, 2, 3, 4, 5)
	fmt.Println("sum_result=", sum_result)
}
