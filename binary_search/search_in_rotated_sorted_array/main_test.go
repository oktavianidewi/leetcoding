package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {
	// nums := []int{6, 7, 0, 1, 2, 4, 5}
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{1, 3}
	// [0, 1, 2, 4, 5, 6, 7]
	target := 4
	expected := -1

	answer := search(nums, target)
	assert.Equal(t, expected, answer)
}
