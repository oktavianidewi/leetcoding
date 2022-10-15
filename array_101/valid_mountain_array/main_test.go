package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validMountainArray(t *testing.T) {
	// t.Skip()
	arr := []int{10, 2, 5, 3}
	expected := false
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_validMountainArray_2(t *testing.T) {
	// t.Skip()
	arr := []int{7, 1, 14, 11}
	expected := false
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_validMountainArray_3(t *testing.T) {
	// t.Skip()
	arr := []int{7, 3, 1, 3, 7}
	expected := false
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_validMountainArray_4(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{-2, 0, 10, -19, 4, 6, -8}
	expected := false
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_validMountainArray_5(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{0, 0}
	expected := false
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_validMountainArray_6(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{0, 3, 2, 1}
	expected := true
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_validMountainArray_7(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{0, 2, 3, 4, 5, 2, 1, 0}
	expected := true
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_validMountainArray_8(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{0, 2, 3, 3, 5, 2, 1, 0}
	expected := false
	answer := validMountainArray(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}
