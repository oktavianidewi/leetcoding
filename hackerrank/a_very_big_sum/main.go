package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	return _sum(ar, len(ar))
}

var res int64

func _sum(ar []int64, index int) int64 {
	if index == 0 {
		return res
	}
	res = res + ar[index-1]
	return _sum(ar, index-1)
}

func main() {
	nums := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	sum := aVeryBigSum(nums)
	fmt.Printf("result sum %v \n", sum)
}
