package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_containsDuplicate_0(t *testing.T) {
	// t.Skip()
	arr := []int{2, 7, 11, 15}
	expected := false
	answer := containsDuplicate(arr)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_containsDuplicate_1(t *testing.T) {
	// t.Skip()
	arr := []int{1, 2, 3, 1}
	expected := true
	answer := containsDuplicate(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_containsDuplicate_2(t *testing.T) {
	// t.Skip()
	arr := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	expected := true
	answer := containsDuplicate(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}
