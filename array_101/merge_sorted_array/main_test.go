package main

import (
	"fmt"
	"testing"
)

func Test_merge(t *testing.T) {
	// t.Skip()
	nums1 := []int{1, 2, 3, 0, 0, 0}
	// [1, 2, 2, 3, 0, 0]
	// [1, 2, 2, 3, 5, 0]
	// [1, 2, 2, 3, 5, 6]
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	// expected := []int{1, 2, 2, 3, 5, 6}
	merge(nums1, m, nums2, n)
	fmt.Printf("xxx %v \n", nums1)
	// assert.Equal(t, expected, answer)
}

func Test_merge_1(t *testing.T) {
	// t.Skip()

	nums1 := []int{1}
	m := 1
	nums2 := []int{}
	n := 0

	// expected := []int{1}
	merge(nums1, m, nums2, n)
	fmt.Printf("xxx %v \n", nums1)

}

func Test_merge_2(t *testing.T) {
	// t.Skip()

	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	m := 6
	nums2 := []int{1, 2, 2}
	n := 3

	// expected := []int{-1,0,0,1,2,2,3,3,3}
	merge(nums1, m, nums2, n)
	fmt.Printf("xxx %v \n", nums1)

}
