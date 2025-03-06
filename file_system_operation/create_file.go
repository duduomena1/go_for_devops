package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./file_system_operation/example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File created successfully")
}
