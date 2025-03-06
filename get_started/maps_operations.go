package main

import (
	"fmt"
)

func main() {
	fmt.Println("Map operations:")
	//initialize a map
	fruits := map[string]int{"apple": 5, "banana": 7}
	//check length
	fmt.Println("length of the map:", len(fruits))
	//adding elements
	fruits["orange"] = 10
	fmt.Println("Added 'orange' to the map:", fruits)

	//retrieving elements
	applePrice := fruits["apple"]
	fmt.Println("Price of apple:", applePrice)
	// checking existence
	price, exist := fruits["kiwi"]
	if exist {
		fmt.Println("Price of kiwi:", price)
	} else {
		fmt.Println("Kiwi not found in the map")
	}

	orangePrice, exist := fruits["orange"]
	if exist {
		fmt.Println("price of orange", orangePrice)
	} else {
		fmt.Println("orange not found in the map")
	}

	// delete an element
	delete(fruits, "banana")
	fmt.Println("After Deleting 'banana':", fruits)

	for key, value := range fruits {
		fmt.Printf("%s: %d\n", key, value)
	}
}
