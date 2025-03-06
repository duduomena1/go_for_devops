package main

import (
	"fmt"
	"os"
)

func main() {
	path := "./file_system_operation/example"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder Created successfully:", path)
	} else {
		fmt.Println("Folder already exists:", path)
	}
}
