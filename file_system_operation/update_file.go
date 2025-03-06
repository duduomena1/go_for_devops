package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./file_system_operation/example.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString("Updated content\n"); err != nil {
		fmt.Println("Error updating file:", err)
		return
	}

	fmt.Println("File updated successfully")
}
