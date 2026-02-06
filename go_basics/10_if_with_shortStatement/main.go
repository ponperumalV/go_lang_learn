package main

import "fmt"

func main() {

	items := 3
	pricePerItem := 42

	if totalPrice := items * pricePerItem; totalPrice > 100 {
		fmt.Println("shipping cost is free")
	} else {
		fmt.Println("shipping cost is 10")
	}
}
