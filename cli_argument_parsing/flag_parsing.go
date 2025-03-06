package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "Output.txt", "The name of the file to write to")
	message := flag.String("message", "Hello, World!", "The message to write to the file")

	flag.Parse()

	file, err := os.OpenFile(*filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(*message + "\n")
	if err != nil {
		fmt.Printf("Error writing to file: %s\n", err)
		return
	}

	fmt.Printf("Message written to %s\n", *filename)
}
