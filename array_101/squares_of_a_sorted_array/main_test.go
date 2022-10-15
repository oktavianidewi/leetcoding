package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedSquares(t *testing.T) {
	// t.Skip()
	nums := []int{-4, -1, 0, 3, 10}
	expected := []int{0, 1, 9, 16, 100}
	answer := sortedSquares(nums)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_sortedSquares_1(t *testing.T) {
	// t.Skip()
	nums := []int{-7, -3, 2, 3, 11}
	expected := []int{4, 9, 9, 49, 121}
	answer := sortedSquares(nums)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_sortedSquares_insert(t *testing.T) {
	t.Skip()
	nums := []int{0, 0, 0}
	expected := []int{0, 10, 0, 0}
	answer := sortedSquares(nums)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}
