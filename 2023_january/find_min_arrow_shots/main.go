package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/solutions/3000437/simple-c-code/?orderBy=newest_to_oldest
	sort.Slice(points, func(a, b int) bool {
		k := len(points[0]) - 1
		return points[a][k] < points[b][k]
	})

	arrow := 1
	last := points[0][1]
	for i := 0; i < len(points); i++ {
		if points[i][0] > last {
			arrow += 1
			last = points[i][1]
		} else {
			last = min(last, points[i][1])
		}
	}
	fmt.Println(points)
	return arrow
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://leetcode.com/discuss/study-guide/2166045/line-sweep-algorithms
func main() {
	points, exp := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4
	// points, exp := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2
	// points, exp := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2
	res := findMinArrowShots(points)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
