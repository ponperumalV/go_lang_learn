package main

import "fmt"

func main() {

	// make ([]t , len , cap)

	score := make([]int, 0, 5)
	fmt.Println(score, len(score), cap(score))

	score1 := make([]int, 3, 5)
	score1[0] = 4
	score1[1] = 5
	score1[2] = 6
	fmt.Println(score1, len(score1), cap(score1))

	score1 = append(score1, 90, 80)
	fmt.Println("after append 90 .80:", score1, len(score1), cap(score1))

	score = append(score, 90, 80)
	fmt.Println("after append 90 .80:", score, len(score), cap(score))

	score = append(score, 70, 60, 50)
	fmt.Println("after append 70,60,50:", score, len(score), cap(score))
	//

	//if we append more than the capacity of the slice, Go will automatically create a new underlying array with a larger capacity and copy the existing elements to it. The new capacity is typically doubled to accommodate future growth.

	score = append(score, 40, 30, 20, 10)
	fmt.Println("after append 40,30,20,10:", score, len(score), cap(score))

	// append multiple slices together

	todo := []string{"task1", "task2", "task3"}

	todo2 := []string{"task5", "task6", "task7"}

	todo = append(todo, todo2...)
	fmt.Println(todo)

}
