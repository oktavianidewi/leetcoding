package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkIfExist(t *testing.T) {
	// t.Skip()
	arr := []int{10, 2, 5, 3}
	expected := true
	answer := checkIfExist(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_checkIfExist_2(t *testing.T) {
	// t.Skip()
	arr := []int{7, 1, 14, 11}
	expected := true
	answer := checkIfExist(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_checkIfExist_3(t *testing.T) {
	// t.Skip()
	arr := []int{3, 1, 7, 11}
	expected := false
	answer := checkIfExist(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_checkIfExist_4(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{-2, 0, 10, -19, 4, 6, -8}
	expected := false
	answer := checkIfExist(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}

func Test_checkIfExist_5(t *testing.T) {
	// kalau hasil bagi dan pembagi memiliki index sama (0), return false
	// t.Skip()
	arr := []int{0, 0}
	expected := true
	answer := checkIfExist(arr)

	fmt.Printf("nums, answer %v \n", answer)
	assert.Equal(t, expected, answer)
}
