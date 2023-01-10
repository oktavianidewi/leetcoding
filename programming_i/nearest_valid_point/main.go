package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/discuss/1096943/Golang-solution-100-100-with-explanation-and-images

func nearestValidPoint(x int, y int, points [][]int) int {
	// return index of the smallest
	minVal := math.MaxInt32
	idx := -1

	for i, point := range points {
		x1, y1 := point[0], point[1]
		if x == x1 && manhattanDistance(y, y1) < minVal {
			minVal = manhattanDistance(y, y1)
			idx = i
		} else if y == y1 && manhattanDistance(x, x1) < minVal {
			minVal = manhattanDistance(x, x1)
			idx = i
		}
	}
	return idx
}

func manhattanDistance(val, val1 int) int {
	return absVal(val - val1)
}

func absVal(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	// x := 3
	// y := 4
	// points := [][]int{{3, 4}}
	// exp := 0

	// points := [][]int{{2, 3}}
	// exp := -1

	// points := [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}
	// exp := 2

	x := 1
	y := 1
	points := [][]int{{1, 2}, {3, 3}, {3, 3}}
	exp := 0

	res := nearestValidPoint(x, y, points)
	fmt.Printf("exp %v, res %v \n", exp, res)
}
