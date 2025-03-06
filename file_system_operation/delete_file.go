package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("./file_system_operation/example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File successfully deleted")

}
