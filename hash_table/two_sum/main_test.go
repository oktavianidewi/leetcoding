package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum_0(t *testing.T) {
	// t.Skip()
	arr := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	answer := twoSum(arr, target)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_twoSum_1(t *testing.T) {
	// t.Skip()
	arr := []int{2, 3, 4}
	target := 6
	expected := []int{0, 2}
	answer := twoSum(arr, target)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	// assert.Equal(t, expected, answer)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_twoSum_2(t *testing.T) {
	// t.Skip()
	arr := []int{-1, 0}
	target := -1
	expected := []int{0, 1}
	answer := twoSum(arr, target)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	// assert.Equal(t, expected, answer)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

// gagal
func Test_twoSum_3(t *testing.T) {
	// t.Skip()
	arr := []int{-3, 3, 4, 90}
	target := 0
	expected := []int{0, 1}
	answer := twoSum(arr, target)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	// assert.Equal(t, expected, answer)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_twoSum_4(t *testing.T) {
	// t.Skip()
	arr := []int{3, 3}
	target := 6
	expected := []int{0, 1}
	answer := twoSum(arr, target)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	// assert.Equal(t, expected, answer)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_twoSum_5(t *testing.T) {
	// t.Skip()
	arr := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	answer := twoSum(arr, target)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	// assert.Equal(t, expected, answer)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}
