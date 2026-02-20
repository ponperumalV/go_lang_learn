package main

import "fmt"

func main() {

	// for range over slice

	views := []int{90, 80, 70, 60, 50}
	total := 0
	mixedVariables := []interface{}{"string", 1, 2.5, true}

	fmt.Println(mixedVariables)

	for i, v := range views {
		fmt.Printf("index: %d , value: %d \n", i, v)
		total = total + v
	}

	fmt.Println("total views:", total)

}
