func orangesRotting(grid [][]int) int {
	var que [][]int
	rotten := 2
	fresh := 1
	freshCounter := 0

	minutes := 0

	// find the rotten
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// fmt.Printf("x %v , y %v, value %v \n", i, j, grid[i][j])
			if grid[i][j] == rotten {
				que = append(que, []int{i, j})
			}

			if grid[i][j] == fresh {
				freshCounter++
			}
		}
	}

	var rotCounter int

	for {
		rotCounter = 0
		var tmp [][]int
		for i := 0; i < len(que); i++ {
			// fmt.Printf("i %v, len que %v \n", i, len(que))
			x := que[i][0]
			y := que[i][1]
			directions := [][]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
			for _, direction := range directions {
				if direction[0] < 0 || direction[0] >= len(grid) || direction[1] < 0 || direction[1] >= len(grid[0]) {
					// fmt.Printf("isi direction0 %v, direction1 %v \n", direction[0], direction[1])
					continue
				}
				// fmt.Printf("not continue\n")
				if grid[direction[0]][direction[1]] == fresh {
					grid[direction[0]][direction[1]] = 2
					rotCounter++
					freshCounter--
					tmp = append(tmp, []int{direction[0], direction[1]})
				}
			}
		}
		if rotCounter == 0 {
			break
		}
		minutes++
		que = tmp
		// fmt.Printf("que %v \n", que)
	}

	if freshCounter != 0 {
		return -1
	}

	return minutes
}