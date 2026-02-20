package main

import (
	"fmt"
)

func main() {

	// map[keyType]valueType{key1: value1, key2: value2, ...}

	student := map[string]int{
		"John":  20,
		"Jane":  22,
		"Steve": 19,
	}

	// fmt.Println(student)
	// // access value by key
	// fmt.Println("John's age:", student["John"])

	// var name string
	// fmt.Print("Enter student name to look up: ")
	// fmt.Scanln(&name)

	// check if a key exists
	age, ok := student["Jane"]
	if ok {
		fmt.Println("Jane's age:", age)
	} else {
		fmt.Println("Jane's age not found")
	}

	// if age, ok := student[name]; ok {
	// 	fmt.Printf("Found it! %s is %d years old.\n", name, age)
	// } else {
	// 	fmt.Printf("Sorry, %s is not in our records.\n", name)
	// }

	// make(map[keyType]valueType)

	var scores map[string]int /* nil map */

	fmt.Println(scores, scores["a"])

	scores = make(map[string]int)

	scores["math"] = 90
	scores["english"] = 80
	fmt.Println(scores)
	fmt.Println("math score:", scores["math"])

	// delete a key-value pair
	delete(scores, "english")
	fmt.Println(scores)

	users := map[string]string{

		"user1": "John",
		"user2": "Jane",
		"user3": "Steve",
	}

	fmt.Println(users)

	delete(users, "user2")
	delete(users, "user4") // no error
	fmt.Println(users)
}
