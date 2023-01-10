package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/minimum-sum-of-four-digit-number-after-splitting-digits/solutions/1750426/golang-solution-with-explanation/
func minimumSum(num int) int {
	var temp []int
	for num > 0 {
		temp = append(temp, num%10)
		num = num / 10
	}
	sort.Ints(temp)
	return (temp[0] * 10) + (temp[1] * 10) + temp[2] + temp[3]
}

func main() {
	num, exp := 2932, 52
	// num, exp := 4009, 13
	res := minimumSum(num)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
