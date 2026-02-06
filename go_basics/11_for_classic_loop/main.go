package main

import "fmt"

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	N := 10
	sum := 0
	for i := 1; i <= N; i++ {
		sum = sum + i
	}
	fmt.Printf("Sum of first %d natural numbers is %d\n", N, sum)
}
