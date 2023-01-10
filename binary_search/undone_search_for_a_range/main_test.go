package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SearchRange_NormalCase(t *testing.T) {
	t.Skip()

	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	expected := []int{3, 4}
	answer := searchRange(nums, target)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_NotFound(t *testing.T) {
	t.Skip()

	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	expected := []int{-1, -1}
	answer := searchRange(nums, target)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_EmptyArray(t *testing.T) {
	t.Skip()

	nums := []int{}
	target := 0
	expected := []int{-1, -1}
	answer := searchRange(nums, target)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_FoundInSingleArray(t *testing.T) {
	t.Skip()

	nums := []int{1}
	target := 1
	expected := []int{0, 0}
	answer := searchRange(nums, target)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_NotFoundInSingleArray(t *testing.T) {
	t.Skip()

	nums := []int{1}
	target := 0
	expected := []int{-1, -1}
	answer := searchRange(nums, target)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_DoubleArray(t *testing.T) {
	t.Skip()

	nums := []int{2, 2}
	target := 2
	expected := []int{0, 1}
	answer := searchRange(nums, target)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_DoubleArrayX(t *testing.T) {

	nums := []int{1, 3}
	target := 1
	expected := []int{0, 0}
	answer := searchRange(nums, target)
	fmt.Printf("DoubleArrayX %v, %v \n", answer, expected)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_DoubleArrayY(t *testing.T) {
	nums := []int{1, 3}
	target := 5
	expected := []int{-1, -1}
	answer := searchRange(nums, target)
	fmt.Printf("DoubleArrayY %v, %v \n", answer, expected)

	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_DoubleArrayZ(t *testing.T) {
	nums := []int{1, 3}
	target := 3
	expected := []int{1, 1}
	answer := searchRange(nums, target)
	fmt.Printf("DoubleArrayZ %v, %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_DoubleArrayA(t *testing.T) {
	nums := []int{2, 2}
	target := 1
	expected := []int{-1, -1}
	answer := searchRange(nums, target)

	fmt.Printf("DoubleArrayA %v, %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}

func Test_SearchRange_DoubleArrayB(t *testing.T) {
	nums := []int{2, 2}
	target := 2
	expected := []int{1, 1}
	answer := searchRange(nums, target)

	fmt.Printf("xxx %v %v \n", answer, expected)
	for i, _ := range expected {
		assert.Equal(t, expected[i], answer[i])
	}
}
