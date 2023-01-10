package main

import "fmt"

func minimumRounds(tasks []int) int {
	// create mapTask
	mapTask := make(map[int]int)
	for _, val := range tasks {
		mapTask[val] += 1
	}
	fmt.Println(mapTask)

	var sumOfCounter int
	for _, task := range mapTask {
		// fmt.Printf("task %v \n", task)

		if task == 1 {
			return -1
		}

		sumOfCounter += task / 3
		if task%3 == 1 {
			sumOfCounter += 1
		}
		if task%3 == 2 {
			sumOfCounter += 1
		}
		// fmt.Printf("counter %v \n", counter)
	}
	return sumOfCounter
}

func main() {
	// tasks, exp := []int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}, 4
	// tasks, exp := []int{2, 3, 3}, -1
	// tasks, exp := []int{5, 5, 5, 5}, 2
	tasks, exp := []int{69, 65, 62, 64, 70, 68, 69, 67, 60, 65, 69, 62, 65, 65, 61, 66, 68, 61, 65, 63, 60, 66, 68, 66, 67, 65, 63, 65, 70, 69, 70, 62, 68, 70, 60, 68, 65, 61, 64, 65, 63, 62, 62, 62, 67, 62, 62, 61, 66, 69}, 20

	res := minimumRounds(tasks)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
