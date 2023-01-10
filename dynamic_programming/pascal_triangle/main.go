package main

import "fmt"

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	triangles := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		triangle := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				triangle[j] = 1
			} else {
				triangle[j] = triangles[i-1][j-1] + triangles[i-1][j]
			}
		}
		triangles[i] = triangle
	}
	// fmt.Printf("isi triangles %v\n", triangles)

	return triangles
}

func main() {
	triangle := generate(5)
	fmt.Printf("triangle %v \n", triangle)
}
