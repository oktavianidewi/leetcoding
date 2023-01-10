package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pivotIndex_1(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{1, 7, 3, 6, 5, 6}
	expected := 3
	answer := pivotIndex(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_pivotIndex_2(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{2, 3, -1, 8, 4}
	expected := 3
	answer := pivotIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_pivotIndex_3(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{1, -1, 4}
	expected := 2
	answer := pivotIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_pivotIndex_4(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{1, 2, 3}
	expected := -1
	answer := pivotIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_pivotIndex_5(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{2, 1, -1}
	expected := 0
	answer := pivotIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}
