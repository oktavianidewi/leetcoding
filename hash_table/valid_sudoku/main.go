package main

import (
	"fmt"
	"time"
)

func isValidSudoku(board [][]byte) bool {
	// make 3 2d array
	rowbuf, colbuf, boxbuf := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	// assign key 1-9 with value boolean (true) to each 2d array
	for i := 0; i < 9; i++ {
		rowbuf[i] = make([]bool, 9)
		colbuf[i] = make([]bool, 9)
		boxbuf[i] = make([]bool, 9)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				// num menentukan index dari angka di sudoku
				num := board[r][c] - '0' - byte(1)

				// 0/3*3 + 0/3 = 0
				fmt.Printf("num %v \n", num)
				fmt.Printf("rowbuf %v \n", rowbuf[r][num])
				fmt.Printf("colbuf %v \n", colbuf[r][num])
				fmt.Printf("boxbuf %v, %v \n", r/3*3+c/3, boxbuf[r/3*3+c/3][num])

				fmt.Printf("hasil %v \n\n", rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num])

				if rowbuf[r][num] || colbuf[c][num] || boxbuf[r/3*3+c/3][num] {
					return false
				}

				rowbuf[r][num] = true
				colbuf[c][num] = true
				boxbuf[r/3*3+c/3][num] = true
			}
			time.Sleep(1 * time.Second)
		}
	}
	return true
}
