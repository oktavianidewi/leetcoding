package main

import (
	"fmt"
)

func minTimeToVisitAllPoints(points [][]int) int {
	minTime := 0
	for i := 0; i < len(points)-1; i++ {
		x := abs(points[i][0] - points[i+1][0])
		y := abs(points[i][1] - points[i+1][1])
		minTime += max(x, y)
	}
	return minTime
}
func abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func main() {
	points, exp := [][]int{{1, 1}, {3, 4}, {-1, 0}}, 7
	// points, exp := [][]int{{3, 2}, {-2, 2}}, 5
	// points, exp := [][]int{{3, 2}, {3, 9}}, 7
	res := minTimeToVisitAllPoints(points)
	fmt.Printf("expected %v, result %v \n", exp, res)

}
