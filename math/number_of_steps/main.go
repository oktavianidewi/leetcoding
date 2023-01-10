package main

import "fmt"

func numberOfSteps(num int) int {
	counter := 0
	for num > 0 {
		if num%2 == 1 {
			num = num - 1
			counter += 1
		} else {
			num = num / 2
			counter += 1
		}
	}
	return counter
}

func main() {
	// num, exp := 14, 6
	num, exp := 123, 12
	// num, exp := 8, 4
	res := numberOfSteps(num)
	fmt.Printf("expected %v, res %v \n", exp, res)
}
