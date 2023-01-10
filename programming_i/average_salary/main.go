package main

import "fmt"

func average(salary []int) float64 {
	var ans float64
	var sum, maxVal, minVal int

	sum = salary[len(salary)-1]
	for i := 0; i < len(salary)-1; i++ {
		sum = sum + salary[i]

		if maxVal == 0 || maxVal < max(salary[i], salary[i+1]) {
			maxVal = max(salary[i], salary[i+1])
		}

		if minVal == 0 || minVal > min(salary[i], salary[i+1]) {
			minVal = min(salary[i], salary[i+1])
		}
	}
	ans = float64(sum-(maxVal+minVal)) / float64(len(salary)-2)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	a := average([]int{4000, 3000, 1000, 2000})
	b := average([]int{1000, 2000, 3000})
	fmt.Printf("a %v \n", a)
	fmt.Printf("b %v \n", b)
}
