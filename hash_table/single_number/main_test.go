package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_singleNumber_0(t *testing.T) {
	// t.Skip()
	arr := []int{2, 2, 1}
	expected := 1
	answer := singleNumber(arr)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_singleNumber_1(t *testing.T) {
	// t.Skip()
	arr := []int{4, 1, 2, 1, 2}
	expected := 4
	answer := singleNumber(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_singleNumber_2(t *testing.T) {
	// t.Skip()
	arr := []int{1}
	expected := 1
	answer := singleNumber(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}
