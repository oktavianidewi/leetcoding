package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	fmt.Printf("initial %d \n", num)
	for num > 0 {
		fmt.Printf(">>> andNum %v \n", num&1)
		// num&1 == 1 sama dengan num%2 == 1
		if num%2 == 1 {
			count++
		}
		num >>= 1 // divided by 2
		// num = num >> 1
		fmt.Printf("increment num value %v \n", num)
	}
	return count
}

func main() {
	var num uint32
	num = 00000000000000000000000000001011
	res := hammingWeight(num)
	fmt.Printf("res %v \n", res)
}
