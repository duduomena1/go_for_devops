package main

import (
	"fmt"
	"regexp"
)

func main() {

	name := "Matheus Jesus"
	Email := "matheusjesus@gmail.com"
	phoneNumber := "+55-71-99999-9999"

	NamePattern := `^[a-zA-Z]+(?: [a-zA-z]+)?$`
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+\d{2}-\d{2}-\d{5}-\d{4}$`

	matchName, _ := regexp.MatchString(NamePattern, name)
	matchEmail, _ := regexp.MatchString(emailPattern, Email)
	matchPhone, _ := regexp.MatchString(phonePattern, phoneNumber)

	fmt.Printf("Name valid: %t\n", matchName)
	fmt.Printf("Email valid: %t\n", matchEmail)
	fmt.Printf("Phone valid: %t\n", matchPhone)
}
