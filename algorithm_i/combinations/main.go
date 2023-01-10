package main

import (
	"fmt"
	"time"
)

func factor(num int) int {
	if num == 1 {
		return 1
	}
	return num * factor(num-1)
}
func combine(n int, k int) [][]int {
	comb := factor(n) / (factor(k) * factor(n-k))
	fmt.Printf("combinatorial %v \n", comb)

	var variations [][]int
	mix := make([]int, k)
	j := 1
	for i := 1; i <= n; {
		mix[0] = i
		fmt.Printf("i %v\n", i)
		for idx := 0; idx < k; {
			fmt.Printf("idx %v\n", idx)
			mix[idx] = j
			fmt.Printf("mix %v\n", mix)

			idx++
		}
		j++
		if j == n {
			i++
			j = i + 1
		}
		time.Sleep(1 * time.Second)
		variations = append(variations, mix)
	}
	fmt.Printf("mix %v \n", mix)
	return variations
}

func main() {
	//
	n, k := 4, 2
	output := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}

	// n, k := 1, 1
	// output := [][]int{{1}}

	result := combine(n, k)
	fmt.Printf("expected %v, result %v \n", output, result)
}
