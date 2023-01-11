package main

import "fmt"

func minTime(n int, edges [][]int, hasApple []bool) int {
	var steps int

	return steps
}

func main() {
	n, edges, hasApple, exp := 7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, true, true, false}, 8
	// n, edges, hasApple, exp := 7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, true, false, false, true, false}, 6
	// n, edges, hasApple, exp := 7, [][]int{{0, 1}, {0, 2}, {1, 4}, {1, 5}, {2, 3}, {2, 6}}, []bool{false, false, false, false, false, false, false}, 0
	res := minTime(n, edges, hasApple)
	fmt.Printf("expected %v, result %v", exp, res)
}
