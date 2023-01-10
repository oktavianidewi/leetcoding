package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindPeakElement(t *testing.T) {
	nums := []int{3, 2, 1, 0}
	expected := 0

	// nums := []int{1, 2, 3}
	// expected := 2

	// nums := []int{1, 2, 3, 1}
	// expected := 2

	// nums := []int{1, 2, 1, 3, 5, 6, 4}
	// expected := 5

	answer := findPeakElement(nums)
	assert.Equal(t, expected, answer)
}
