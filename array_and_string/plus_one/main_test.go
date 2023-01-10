package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_plusOne_0(t *testing.T) {
	t.Skip()
	arr := []int{9}
	expected := []int{1, 0}
	answer := plusOne(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_plusOne_1(t *testing.T) {
	// t.Skip()
	arr := []int{9, 9}
	expected := []int{1, 0, 0}
	answer := plusOne(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}
