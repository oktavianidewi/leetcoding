package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_duplicateZeros(t *testing.T) {
	// t.Skip()
	nums := []int{1, 0, 2, 3, 0, 4, 5, 0}
	expected := []int{1, 0, 0, 2, 3, 0, 0, 4}
	duplicateZeros(nums)

	fmt.Printf("nums, answer %v \n", nums)
	for i, _ := range expected {
		assert.Equal(t, expected[i], nums[i])
	}
}

func Test_duplicateZeros_1(t *testing.T) {
	// t.Skip()
	nums := []int{1, 2, 3}
	expected := []int{1, 2, 3}
	duplicateZeros(nums)

	fmt.Printf("nums, answer %v \n", nums)
	for i, _ := range expected {
		assert.Equal(t, expected[i], nums[i])
	}
}

func Test_duplicateZeros_2(t *testing.T) {
	// t.Skip()
	nums := []int{1, 0, 0, 2, 3, 0, 0}
	expected := []int{1, 0, 0, 0, 0, 2, 3}
	duplicateZeros(nums)

	fmt.Printf("nums, answer %v \n", nums)
	for i, _ := range expected {
		assert.Equal(t, expected[i], nums[i])
	}
}

func Test_duplicateZeros_3(t *testing.T) {
	// // t.Skip()
	nums := []int{0, 0, 0, 0, 0, 0, 0}
	expected := []int{0, 0, 0, 0, 0, 0, 0}
	duplicateZeros(nums)

	fmt.Printf("nums, answer %v, %v \n", nums, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], nums[i])
	}
}
