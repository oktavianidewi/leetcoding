package main

import (
	"fmt"
)

func maximum69Number(num int) int {
	// list all possible numbers, record changing max
	maxVal := num
	ori := num
	counter := 1
	for num > 0 {
		var before, after int
		before = (num % 10) * counter
		if (num % 10) == 9 {
			after = 6 * counter
		} else {
			after = 9 * counter
		}
		newVal := ori - before + after
		if maxVal < newVal {
			maxVal = newVal
		}
		counter *= 10
		num /= 10

	}
	return maxVal
}

func main() {
	num, exp := 9999, 9999
	// num, exp := 9669, 9969
	res := maximum69Number(num)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
