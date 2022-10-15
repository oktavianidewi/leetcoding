package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDiagonalOrder_1(t *testing.T) {
	// t.Skip()
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}
	answer := findDiagonalOrder(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_findDiagonalOrder_2(t *testing.T) {
	// t.Skip()
	arr := [][]int{{1, 2}, {3, 4}}
	expected := []int{1, 2, 3, 4}
	answer := findDiagonalOrder(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}
