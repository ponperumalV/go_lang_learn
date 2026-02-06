package main

import "fmt"

func main() {

	isLoggedIn := true
	isAdmin := false
	isSubscribed := true

	canDeletePost := isAdmin || isLoggedIn
	canViewContent := isLoggedIn && isSubscribed

	fmt.Println("Can Delete Post:", canDeletePost)
	fmt.Println("Can View Content:", canViewContent)

	age := 25
	isAdult := age >= 18
	isSenior := age >= 65
	fmt.Println("Is Adult:", isAdult)
	fmt.Println("Is Senior:", isSenior)
}
