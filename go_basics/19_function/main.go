package main

import "fmt"

func greet() {
	fmt.Println("Hello, welcome to Go programming!")
}

func sum(a int, b int) int {
	return a + b
}

func sumAndMultiply(a int, b int) (int, int) {
	sum := a + b
	product := a * b
	return sum, product
}

func main() {
	// function with no parameters and no return value
	greet()

	// function with parameters and a return value
	result := sum(5, 3)
	fmt.Println("Sum:", result)
	// function with multiple parameters
	result2 := sum(10, 20)
	fmt.Println("Sum:", result2)

	// function with multiple return values
	sumResult, productResult := sumAndMultiply(5, 3)
	fmt.Println("Sum:", sumResult, "Product:", productResult)

	// _ blank identifier to ignore a return value

	_, productOnly := sumAndMultiply(7, 4)
	fmt.Println("Product Only:", productOnly)
}
