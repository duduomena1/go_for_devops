package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("./file_system_operation/example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:\n", string(data))
}
