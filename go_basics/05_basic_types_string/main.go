package main

import (
	"fmt"
	"strings"
)

func main() {

	firstName := "Ponperumal"
	lastName := "Vp"
	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)
	fmt.Println("Capitalized Full Name:", strings.ToUpper(fullName))
}
