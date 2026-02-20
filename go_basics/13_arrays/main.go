package main

import "fmt"

func main() {

	// fixed and can not grow

	var mark [5]int

	mark[0] = 90
	mark[1] = 80
	mark[2] = 70
	mark[3] = 60
	mark[4] = 50

	fmt.Println(mark)

	// array literal

	mark2 := [5]int{90, 80, 70, 60, 50}

	fmt.Println(len(mark2))
}
