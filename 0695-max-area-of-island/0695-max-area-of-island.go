
func maxAreaOfIsland(grid [][]int) int {
	max := 0
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			area := calculateArea(grid, x, y)
			if max < area {
				max = area
			}
		}
	}
	return max
}

func calculateArea(matrices [][]int, x, y int) int {
	area := 0

	// avoid boundaries
	if x >= len(matrices) || x < 0 || y >= len(matrices[0]) || y < 0 || matrices[x][y] == 0 {
		return area
	}

	// set zero supaya tidak endless loop
	matrices[x][y] = 0
	area = 1

	area += calculateArea(matrices, x+1, y) + calculateArea(matrices, x-1, y) + calculateArea(matrices, x, y+1) + calculateArea(matrices, x, y-1)

	return area
}