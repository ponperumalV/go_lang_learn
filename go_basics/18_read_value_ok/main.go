package main

import "fmt"

func main() {

	points := map[string]int{
		"a": 10,
		"b": 0, //VALID VALUE

	}

	fmt.Println("a:", points["a"]) // 10
	fmt.Println("b:", points["b"]) // 0
	fmt.Println("c:", points["c"]) // 0 (default value for int)

	// read value with ok idiom
	if value, ok := points["a"]; ok {
		fmt.Println("Value for 'a':", value)
	} else {
		fmt.Println("'a' not found")
	}

	if value, ok := points["c"]; ok {
		fmt.Println("Value for 'c':", value)
	} else {
		fmt.Println("'c' key is not found")
	}

	// for loop in map

	prices := map[string]int{

		"apple":  100,
		"banana": 10,
		"orange": 50,
	}

	totalPrice := 0

	for item, price := range prices {
		fmt.Println("Item:", item, "Price:", price)
		totalPrice = totalPrice + price
	}
	fmt.Println("Total Price:", totalPrice)

	for item := range prices {
		fmt.Println("Item:", item)

	}
}
