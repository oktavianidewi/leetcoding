package main

import (
	"fmt"
)

func diagonalDifference(arr [][]int32) int32 {
	// left diagonal (x,y) x up y up, x down y down
	var sumLeftDiagonal int32
	for x := 0; x < len(arr); x++ {
		sumLeftDiagonal = sumLeftDiagonal + arr[x][x]
	}

	// right diagonal (x,y) x up y down, x down y up
	var sumRightDiagonal int32
	y := len(arr) - 1
	for x := 0; x < len(arr); x++ {
		sumRightDiagonal = sumRightDiagonal + arr[x][y]
		y--
	}

	result := sumLeftDiagonal - sumRightDiagonal

	if result < 0 {
		return result * -1
	}
	return result
}

func main() {
	arr := [][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}
	result := diagonalDifference(arr)
	fmt.Printf("result %v \n", result)

}
