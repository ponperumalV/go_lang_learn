package main

import (
	"fmt"
)

func main() {

	amount1 := 1000
	amount2 := 500

	totalAmount := amount1 + amount2

	fmt.Println("Total Amount:", totalAmount)

	person := 1
	person++
	person += 2

	discount := 10.5

	fmt.Println(amount1, amount2)
	fmt.Println("Person Count:", person)
	fmt.Println("Discount:", discount)
}
