package main

import (
	"fmt"
	"testing"
)

func Test_rotate_0(t *testing.T) {
	// t.Skip()
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, k)
	fmt.Printf("nums, answer %v, %v \n", nums, expected)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_rotate_1(t *testing.T) {
	// t.Skip()
	nums := []int{-1, -100, 3, 99}
	k := 2
	expected := []int{3, 99, -1, -100}
	rotate(nums, k)
	fmt.Printf("nums, answer %v, %v \n", nums, expected)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_rotate_2(t *testing.T) {
	// t.Skip()
	nums := []int{1, 2}
	k := 3
	expected := []int{2, 1}
	rotate(nums, k)
	fmt.Printf("nums, answer %v, %v \n", nums, expected)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}
