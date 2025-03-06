package main

import (
	"fmt"
)

func main() {

	var map1 map[string]int
	fmt.Println("Map declared with var keyword without initialization:", map1)

	map2 := map[string]int{"apple": 2, "banana": 6}
	fmt.Println("Map declared with a map literal:", map2)

	map3 := make(map[string]int)
	map3["apple"] = 25
	map3["banana"] = 50
	fmt.Println("Map declared with make function:", map3)
	fmt.Println("Length of map 3:", len(map3))

	map4 := make(map[string]int, 10)
	map4["apple"] = 54
	map4["banana"] = 23
	fmt.Println("Map declared with make function and specific size:", map4)
	fmt.Println("Length of map 4:", len(map4))

	fmt.Println("Using struct as map values:")
	type Fruit struct {
		Price int
		Color string
	}
	map5 := make(map[string]Fruit)
	map5["apple"] = Fruit{Price: 25, Color: "red"}
	map5["banana"] = Fruit{Price: 50, Color: "yellow"}
	map5["kiwi"] = Fruit{Price: 75, Color: "green"}
	fmt.Println("Map with struct values:", map5)

	fmt.Println("Nested maps")
	map6 := make(map[string]map[string]int)
	map6["fruit"] = map[string]int{"apple": 25, "banana": 50}
	map6["vegetable"] = map[string]int{"carrot": 10, "beans": 20}
	fmt.Println("Nested map:", map6)

}
