package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	// nums := []int{3, 4, 5, 1, 2}
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{1, 3}
	// [0, 1, 2, 4, 5, 6, 7]
	expected := 1

	answer := findMin(nums)
	assert.Equal(t, expected, answer)
}
