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