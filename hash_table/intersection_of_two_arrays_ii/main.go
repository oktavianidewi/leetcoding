package main

import "fmt"

// differenciate the index

func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	var res []int
	for _, n := range nums1 {
		m[n]++
	}
	fmt.Printf("hasil m %v \n", m)

	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			// decreament hasil occurence
			m[n]--
		}
	}
	return res
}
