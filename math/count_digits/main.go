package main

import "fmt"

func countDigits(num int) int {
	ori := num
	counter := 0
	for num > 0 {
		if ori%(num%10) == 0 {
			counter += 1
		}
		// fmt.Println(ori, num, num%10)
		num /= 10
	}
	return counter
}

func main() {
	num, exp := 7, 1
	// num, exp := 121, 2
	// num, exp := 1248, 4
	res := countDigits(num)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
