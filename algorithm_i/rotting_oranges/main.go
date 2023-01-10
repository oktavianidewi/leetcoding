package main

import "fmt"

// cari dimana rotten orange
// 4-directional from the rotten orange
// cari dimana fresh orange, dengan direction kanan-kiri, atas-bawah
// jika ada fresh orange: masukkan koordinat tersebut ke sebuah queue
// hitung counter
// jika tidak ada fresh orange atau ketemu rotten orange, continue
// jika semua koordinat di queue sudah habis, return counter

// tidak ada: return -1

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
func orangesRotting_dewi(grid [][]int) int {
	// masih ga bisa2, coba lihat solusi
	// https://leetcode.com/problems/rotting-oranges/discuss/239339/Golang-BFS

	var que [][]int
	rotten := 2
	fresh := 1

	minutes := 0
	freshCounter := 0

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

	if freshCounter == 0 && len(que) == 0 {
		return 0
	} else if freshCounter != 0 && len(que) == 0 {
		return -1
	}

	historyFreshCounter := freshCounter

	for len(que) > 0 {
		x := que[0][0]
		y := que[0][1]

		if x-1 != -1 && grid[x-1][y] == fresh {
			grid[x-1][y] = rotten
			que = append(que, []int{x - 1, y})
			freshCounter--
		}
		if x+1 != len(grid) && grid[x+1][y] == fresh {
			grid[x+1][y] = rotten
			que = append(que, []int{x + 1, y})
			freshCounter--
		}
		if y-1 != -1 && grid[x][y-1] == fresh {
			grid[x][y-1] = rotten
			que = append(que, []int{x, y - 1})
			freshCounter--
		}
		if y+1 != len(grid[0]) && grid[x][y+1] == fresh {
			grid[x][y+1] = rotten
			que = append(que, []int{x, y + 1})
			freshCounter--
		}
		if historyFreshCounter > freshCounter {
			minutes++
		}
		fmt.Printf("history: %v, current: %v, minutes %v \n", historyFreshCounter, freshCounter, minutes)

		historyFreshCounter = freshCounter
		que = que[1:]
		fmt.Printf("coordinate %v, grid %v \n", fmt.Sprint(x, y), grid)
	}

	if minutes == 0 {
		return 0
	}

	fmt.Printf("que %v, counter %v \n", que, minutes)
	return minutes
}

func main() {
	// grid, output := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4
	grid, output := [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1
	// grid, output := [][]int{{2, 1, 1}, {1, 1, 0}, {1, 0, 1}}, -1
	// grid, output := [][]int{{0, 2}}, 0
	// grid, output := [][]int{{1}}, -1
	// grid, output := [][]int{{1, 2}}, 1

	result := orangesRotting(grid)
	fmt.Printf("expected %v, result %v \n", output, result)
}
