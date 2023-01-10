package main

import (
	"fmt"
	"sort"
)

// failed
func maxPoints_dewi(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	var res int
	// var x_diff []int
	y_diff := make([]int, 1)
	for i := 0; i < len(points)-1; i++ {
		// fmt.Printf("x %v \n", points[i][0])
		// x_diff = append(x_diff, points[i][0]-points[i+1][0])

		// fmt.Printf("y %v \n", points[i][1])
		y_diff = append(y_diff, abs(points[i][1]-points[i+1][1]))

		if y_diff[i]-y_diff[i+1] == 0 && res < points[i][1] {
			res = points[i][1]
		}
	}
	fmt.Println(y_diff)
	fmt.Println(points)
	return res
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

// https://leetcode.com/problems/max-points-on-a-line/solutions/1734310/golang-solution-with-illustration/?q=golang&orderBy=most_relevant
func maxPoints(points [][]int) int {
	max := 0

	for i := 0; i < len(points); i++ {
		m := make(map[float64]int) // slope : number of points on the slope
		for j := 0; j < len(points); j++ {
			x1, x2, y1, y2 := points[i][0], points[j][0], points[i][1], points[j][1]

			slope := float64(y2-y1) / float64(x2-x1)
			m[slope]++
		}

		for _, b := range m {
			newB := (b + 1)
			if newB > max {
				max = newB
			}
		}
	}

	if len(points) == 1 {
		return 1
	}

	return max
}

func main() {
	points, exp := [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}, 4
	// points, exp := [][]int{{1, 1}, {2, 2}, {3, 3}}, 3
	res := maxPoints(points)
	fmt.Printf("expected %v, res %v \n", exp, res)
}
