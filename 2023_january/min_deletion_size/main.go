package main

import (
	"fmt"
)

func minDeletionSize(strs []string) int {
	count := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 0; j < len(strs)-1; j++ {
			fmt.Println(string(strs[j][i]), strs[j][i], string(strs[j+1][i]), strs[j+1][i])
			if strs[j][i] > strs[j+1][i] {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	strs, exp := []string{"abc", "def", "fhg"}, 1
	// strs, exp := []string{"abc", "bca", "cee"}, 1
	// strs, exp := []string{"cba", "daf", "ghi"}, 1
	// strs, exp := []string{"zyx", "wvu", "tsr"}, 3
	// strs, exp := []string{"a", "b"}, 0
	res := minDeletionSize(strs)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
