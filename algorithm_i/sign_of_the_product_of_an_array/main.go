package main

import "fmt"

func arraySign(nums []int) int {
	// tidak perlu di-product, tapi jika ada sejumlah -1 (ganjil), pasti hasil -1
	otr := 0
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			otr += 1
		}
	}
	if otr%2 == 0 {
		return 1
	}
	return -1
}

func arraySign_dewi(nums []int) int {
	product := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		} else if num < 0 {
			num = -1
		} else {
			num = 1
		}
		product = product * num
	}

	if product < 0 {
		return -1
	}

	return 1
}
func main() {
	// nums := []int{-1, -2, -3, -4, 3, 2, 1}
	// exp := 1

	// nums := []int{1, 5, 0, 2, -3}
	// exp := 0

	// nums := []int{-1, 1, -1, 1, -1}
	// exp := -1

	nums := []int{9, 72, 34, 29, -49, -22, -77, -17, -66, -75, -44, -30, -24}
	exp := -1

	ans := arraySign(nums)
	fmt.Printf("expected %v, answer %v \n", exp, ans)

}
