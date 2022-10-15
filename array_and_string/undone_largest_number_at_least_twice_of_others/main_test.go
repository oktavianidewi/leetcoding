package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dominantIndex_0(t *testing.T) {
	// t.Skip()
	arr := []int{3, 6, 1, 0}
	expected := 1
	answer := dominantIndex(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_dominantIndex_1(t *testing.T) {
	// t.Skip()
	arr := []int{1, 7, 3, 6, 5}
	expected := -1
	answer := dominantIndex(arr)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_dominantIndex_2(t *testing.T) {
	// t.Skip()
	arr := []int{2, 3, -1, 8, 4}
	expected := 3
	answer := dominantIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_dominantIndex_3(t *testing.T) {
	// t.Skip()
	arr := []int{1, -1, 4}
	expected := 2
	answer := dominantIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_dominantIndex_4(t *testing.T) {
	// t.Skip()
	arr := []int{1, 2, 3}
	expected := -1
	answer := dominantIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_dominantIndex_5(t *testing.T) {
	// t.Skip()
	arr := []int{2}
	expected := 0
	answer := dominantIndex(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}
