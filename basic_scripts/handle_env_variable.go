package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("ENV_VAR", "Hey you")

	myVar := os.Getenv("ENV_VAR")
	if myVar == "" {
		fmt.Println("ENV_VAR is not set")
	} else {
		fmt.Println("ENV_VAR:", myVar)
	}

	//Listing all environment variables
	fmt.Println("All Environment Variables:")
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
}
