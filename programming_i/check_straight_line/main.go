package main

import "fmt"

func checkStraightLine(coordinates [][]int) bool {
	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[1][0], coordinates[1][1]

	for i := 2; i < len(coordinates); i++ {
		x, y := coordinates[i][0], coordinates[i][1]

		// fmt.Printf("result %v, %v \n", (y2-y1)*(x-x1), (x2-x1)*(y-y1))
		if (y2-y1)*(x-x1) != (x2-x1)*(y-y1) {
			return false
		}
	}
	return true
}

func main() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	output := true

	result := checkStraightLine(coordinates)
	fmt.Printf("expected %v, result %v \n", output, result)
}
