package main

import "fmt"

func sunAll(nums ...int) int {
	total := 0

	for _, num := range nums {
		fmt.Println("Adding:", num)
		total += num
	}
	return total
}

// for higher-order function example
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Total Sum:", sunAll(1, 2, 3, 4, 5))

	// passing a slice to variadic function
	numbers := []int{10, 20, 30}
	fmt.Println("Total Sum from slice:", sunAll(numbers...))

	result := func(a int, b int) int {
		return a + b
	}
	fmt.Println("Anonymous Function Result (Sum):", result(5, 3))

	// IIFE (Immediately Invoked Function Expression)
	result1 := func(a int, b int) int {
		return a * b
	}(5, 3)
	fmt.Println("IIFE Result (Product):", result1)

	// higher-order function example
	higherOrderFunc := func(operation func(int, int) int, a int, b int) int {
		return operation(a, b)
	}

	sumResult := higherOrderFunc(add, 10, 20)
	fmt.Println("Higher-Order Function Result (Sum):", sumResult)

}
