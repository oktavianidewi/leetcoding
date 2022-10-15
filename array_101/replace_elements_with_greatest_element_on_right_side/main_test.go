package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replaceElements_1(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{17, 18, 5, 4, 6, 1}
	expected := []int{18, 6, 6, 6, 1, -1}
	answer := replaceElements(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_replaceElements_2(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{400}
	expected := []int{-1}
	answer := replaceElements(arr)

	fmt.Printf("nums, answer %v \n", answer)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}
