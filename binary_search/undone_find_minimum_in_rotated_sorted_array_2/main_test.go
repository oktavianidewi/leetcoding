package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CaseSorted_1(t *testing.T) {
	t.Skip()

	input := []int{1, 3, 5}
	expected := 1

	fmt.Println("Test_CaseSorted_1")
	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_1(t *testing.T) {
	t.Skip()

	input := []int{2, 2, 2, 0, 1}
	expected := 0

	fmt.Println("Test_CaseSortRotated_1")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_2(t *testing.T) {
	t.Skip()

	input := []int{4, 5, 5, 6, 0, 1, 2, 3, 3, 4}
	expected := 0
	fmt.Println("Test_CaseSortRotated_2")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_3(t *testing.T) {
	t.Skip()

	input := []int{3, 3, 4, 4, 5, 5, 6, 0, 1, 2}
	expected := 0
	fmt.Println("Test_CaseSortRotated_3")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_4(t *testing.T) {
	t.Skip()
	input := []int{5, 6, 0, 1, 2, 3, 3, 4, 4, 5}
	expected := 0
	fmt.Println("Test_CaseSortRotated_4")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_5(t *testing.T) {
	t.Skip()
	input := []int{1}
	expected := 1
	fmt.Println("Test_CaseSortRotated_5")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_6(t *testing.T) {
	t.Skip()

	input := []int{1, 2}
	expected := 1
	fmt.Println("Test_CaseSortRotated_6")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_7(t *testing.T) {
	t.Skip()

	input := []int{2, 2}
	expected := 2
	fmt.Println("Test_CaseSortRotated_7")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_8(t *testing.T) {
	t.Skip()

	input := []int{2, 2, 2}
	expected := 2
	fmt.Println("Test_CaseSortRotated_8")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_11(t *testing.T) {
	t.Skip()

	input := []int{10, 10, 10, 1, 10}
	expected := 1
	fmt.Println("Test_CaseSortRotated_9")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_10(t *testing.T) {
	t.Skip()

	input := []int{10, 1, 10, 10, 10}
	expected := 1
	fmt.Println("Test_CaseSortRotated_10")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}
func Test_CaseSortRotated_12(t *testing.T) {
	t.Skip()

	input := []int{3, 3, 3, 3, 3, 3, 3, 3, 1, 3}
	expected := 1

	fmt.Println("Test_CaseSortRotated_12")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_9(t *testing.T) {
	t.Skip()

	input := []int{3, 3, 1, 3}
	expected := 1

	fmt.Println("Test_CaseSortRotated_9")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}

func Test_CaseSortRotated_13(t *testing.T) {
	// t.Skip()

	input := []int{-1, -1, -1, -1}
	expected := 1

	fmt.Println("Test_CaseSortRotated_13")

	answer := findMin(input)
	assert.Equal(t, expected, answer)
}
