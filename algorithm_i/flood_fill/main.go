package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	var que [][]int
	que = append(que, []int{sr, sc})
	startColor := image[sr][sc]

	if startColor == color {
		return image
	}

	for len(que) > 0 {
		// cek neighborhood
		x := que[0][0]
		y := que[0][1]
		image[x][y] = color

		// add to que
		// fmt.Printf("koordinat x %v, value %v \n", x-1, image[x-1][y])
		if x-1 != -1 && image[x-1][y] == startColor {
			image[x-1][y] = color
			que = append(que, []int{x - 1, y})
		}
		if x+1 != len(image) && image[x+1][y] == startColor {
			image[x+1][y] = color
			que = append(que, []int{x + 1, y})
		}
		if y-1 != -1 && image[x][y-1] == startColor {
			image[x][y-1] = color
			que = append(que, []int{x, y - 1})
		}
		if y+1 != len(image[0]) && image[x][y+1] == startColor {
			image[x][y+1] = color
			que = append(que, []int{x, y + 1})
		}

		// fmt.Printf("que %v \n", que)
		// pop
		que = que[1:]
	}
	return image
}

func main() {
	// image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	// sr, sc, color := 1, 1, 2
	// output := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}

	// image := [][]int{{0, 0, 0}, {0, 0, 0}}
	// sr, sc, color := 1, 0, 2
	// output := [][]int{{2, 2, 2}, {2, 2, 2}}

	// image := [][]int{{0, 0, 0}, {0, 0, 0}}
	// sr, sc, color := 0, 0, 0
	// output := [][]int{{0, 0, 0}, {0, 0, 0}}

	image := [][]int{{0, 0, 0}, {0, 1, 0}}
	sr, sc, color := 1, 1, 2
	output := [][]int{{0, 0, 0}, {0, 2, 0}}

	result := floodFill(image, sr, sc, color)
	fmt.Printf("expected %v, result %v \n", output, result)
}
