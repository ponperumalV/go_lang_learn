package main

import "fmt"

func main() {
	// dynamic and can grow
	// []type{value1, value2, ...}
	// Arrays in Go are fixed-size collections of elements of the same type.
	// They have a length that is determined at compile time and cannot be changed.

	mark := []int{90, 80, 70, 60, 50}
	fmt.Println(mark, mark[0], mark[1], mark[len(mark)-1])
	mark[0] = 100
	fmt.Println(mark)

	mark = append(mark, 40)
	fmt.Println(mark)
}
