package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	t.Skip()
	nums := []int{3, 2, 2, 3}
	val := 3
	expected := 2

	// expected := []int{1, 2, 2, 3, 5, 6}
	answer := removeElement(nums, val)
	fmt.Printf("Test_removeElement %v \n", nums)
	assert.Equal(t, expected, answer)
}

func Test_removeElement_1(t *testing.T) {
	// t.Skip()
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expected := 5

	// expected := []int{1, 2, 2, 3, 5, 6}
	answer := removeElement(nums, val)
	fmt.Printf("Test_removeElement %v \n", nums)
	assert.Equal(t, expected, answer)
}
