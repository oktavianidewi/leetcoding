package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_arrayPairSum_0(t *testing.T) {
	// t.Skip()
	arr := []int{1, 4, 3, 2}
	expected := 4
	answer := arrayPairSum(arr)
	fmt.Printf("nums, answer %v, %v \n", arr, expected)
	assert.Equal(t, expected, answer)

	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_arrayPairSum_1(t *testing.T) {
	// t.Skip()
	arr := []int{6, 2, 6, 5, 1, 2}
	expected := 9
	answer := arrayPairSum(arr)

	fmt.Printf("nums, answer %v, %v \n", arr, expected)
	assert.Equal(t, expected, answer)

	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}
