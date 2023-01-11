package main

import "fmt"

func isSelfDivided(i int) bool {
	for j := i; j > 0; j /= 10 {
		k := j % 10
		// fmt.Println(k)
		if k == 0 || i%k != 0 {
			return false
		}
		// fmt.Println(j, i%k, k)
	}
	return true
}
func selfDividingNumbers(left int, right int) []int {
	var res []int
	for i := left; i <= right; i++ {
		if i%10 > 0 {
			if isSelfDivided(i) {
				res = append(res, i)
			}
		}
	}
	return res
}

func main() {
	// left, right := 1, 22
	// exp := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}

	// left, right := 47, 85
	// exp := []int{48, 55, 66, 77}

	left, right := 601, 611
	exp := []int{611}

	res := selfDividingNumbers(left, right)
	fmt.Printf("expected %v, result %v \n", exp, res)
}
