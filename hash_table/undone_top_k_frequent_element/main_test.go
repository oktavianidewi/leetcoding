package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// create test cases
func TestTopKFrequent(t *testing.T) {
	t.Skip()
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	expected := []int{1, 2}

	actual := topKFrequent(nums, k)
	fmt.Printf("%v, %v, \n", expected, actual)
	assert.Equal(t, expected, actual)
}

func TestTopKFrequent_1(t *testing.T) {
	// t.Skip()
	nums := []int{1, 1, 1, 2, 3, 3}
	k := 2
	expected := []int{1, 3}

	actual := topKFrequent(nums, k)
	fmt.Printf("%v, %v, \n", expected, actual)
	assert.Equal(t, expected, actual)
}

func TestTopKFrequent_2(t *testing.T) {
	t.Skip()

	nums := []int{1}
	k := 1
	expected := []int{1}

	actual := topKFrequent(nums, k)
	fmt.Printf("%v, %v, \n", expected, actual)

	assert.Equal(t, expected, actual)
}
