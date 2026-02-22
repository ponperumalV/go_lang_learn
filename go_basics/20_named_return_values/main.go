package main

import "fmt"

func calculate(a int, b int) (apple int, orange int) {

	apple = a + b
	orange = a * b
	return // return with named return values
}

func main() {
	sum, product := calculate(5, 3)
	fmt.Println("Sum (Apple):", sum)
	fmt.Println("Product (Orange):", product)
}
