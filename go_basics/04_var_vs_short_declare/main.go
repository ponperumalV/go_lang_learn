package main

import "fmt"

func main() {

	var city string
	city = "Tirunelveli"

	var village = "ettery" // inferred to string

	// :=
	Amount := 1000 // inferred to int

	Amount = Amount + 500

	totalAmount, balance := 1500, 500

	fmt.Println("City:", city)
	fmt.Println("Village:", village)
	fmt.Println("Amount:", Amount)
	fmt.Println("Total Amount:", totalAmount)
	fmt.Println("Balance:", balance)

}
