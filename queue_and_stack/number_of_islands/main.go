package main

import "fmt"

func numIslands(grid [][]byte) int {
	var queue [][]int
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				queue = append(queue, []int{i, j})
				for len(queue) > 0 {
					row, col := queue[0][0], queue[0][1]
					queue = queue[1:]
					if row > 0 && grid[row-1][col] == '1' {
						queue = append(queue, []int{row - 1, col})
						grid[row-1][col] = '0'
					}
					if row < len(grid)-1 && grid[row+1][col] == '1' {
						queue = append(queue, []int{row + 1, col})
						grid[row+1][col] = '0'
					}
					if col > 0 && grid[row][col-1] == '1' {
						queue = append(queue, []int{row, col - 1})
						grid[row][col-1] = '0'
					}
					if col < len(grid[i])-1 && grid[row][col+1] == '1' {
						queue = append(queue, []int{row, col + 1})
						grid[row][col+1] = '0'
					}
				}
				res++
			}
		}
	}
	return res
}

func numIslands_dewi(grid [][]byte) int {
	var queue [][]int
	res := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			// fmt.Printf("koordinat x, y : (%v, %v), value %v \n", x, y, grid[x][y])
			if grid[x][y] == 1 {
				queue = append(queue, []int{x, y})

				for len(queue) > 0 {
					// cek 4 arah mata angin
					row, col := queue[0][0], queue[0][1]

					// decreament, keluarkan queue paling awal
					queue = queue[1:]

					// cek kiri
					if row > 0 && grid[row-1][col] == 1 {
						queue = append(queue, []int{row - 1, col})
						// kenapa harus di-replace? supaya jelas mana pulau yang udah visited/belum; shg ga memenuhi kondisi grid[row-1][col] == 1
						grid[row-1][col] = '0'
					}
					// cek kanan
					if row < len(grid)-1 && grid[row+1][col] == 1 {
						queue = append(queue, []int{row + 1, col})
						grid[row+1][col] = '0'
					}
					// cek bawah
					if col > 0 && grid[row][col-1] == 1 {
						queue = append(queue, []int{row, col - 1})
						grid[row][col-1] = '0'
					}
					// cek atas
					if col < len(grid[0])-1 && grid[row][col+1] == 1 {
						queue = append(queue, []int{row, col + 1})
						grid[row][col+1] = '0'
					}
				}
				// fmt.Printf(">>> queue %v \n", queue)
				// incerement jika nilai queue kosong
				res++
			}
		}
	}
	// fmt.Printf("grid %v \n", grid)
	// fmt.Printf("res %v \n", res)
	return res
}

// TODO: why different???

func main() {
	// grid := [][]byte{
	// 	{1, 1, 1, 1, 0},
	// 	{1, 1, 0, 1, 0},
	// 	{1, 1, 0, 0, 0},
	// 	{0, 0, 0, 0, 0},
	// }
	// ans := 1

	// grid := [][]byte{
	// 	{1, 1, 0, 1},
	// 	{0, 0, 1, 0},
	// 	{0, 0, 0, 0},
	// }
	// ans := 3

	grid := [][]byte{
		{"1", "1", "1", "1", "0"},
		{"1", "1", "0", "1", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "0", "0", "0"},
	}
	ans := 1

	res := numIslands(grid)
	fmt.Printf("num of islands %v , expected ans %v \n", res, ans)
}
