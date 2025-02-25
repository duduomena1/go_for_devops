package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "hey"
	str2 := "you"

	concatenated := str1 + " " + str2
	fmt.Println("Concatenated:", concatenated)

	//split a strings
	sentence := "I am a strings"
	words := strings.Split(sentence, " ")
	fmt.Println("Words:", words)

	//replace a substrings

	replaced := strings.Replace(sentence, "strings", "gopher", 1)
	fmt.Println("Replaced:", replaced)

	// upper and lower case
	upper := strings.ToUpper(sentence)
	lower := strings.ToLower(sentence)
	fmt.Println("Upper:", upper)
	fmt.Println("Lower:", lower)

	//trimming

	spaceystrings := "     too much space     "
	trimmed := strings.TrimSpace(spaceystrings)
	fmt.Println("Trimmed:", trimmed)

	//substrings -  using slicing
	// Note: Go does not have a built-in 'substrings' method, but we can achieve this with slicing

	start := 5
	end := 8
	if end <= len(sentence) && start < end {
		substrings := sentence[start:end]
		fmt.Println("Substrings:", substrings)
	} else {
		fmt.Println("invalid range for substrings")
	}

	//checking if a strings contains a substrings
	contains := strings.Contains(sentence, "am")
	fmt.Println("Contains 'am':", contains)

	index := strings.Index(sentence, "am")
	fmt.Println("Index of 'am':", index)

}
