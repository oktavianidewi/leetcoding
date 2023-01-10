package main

import (
	"fmt"
	"testing"
)

func Test_reverseString_0(t *testing.T) {
	// t.Skip()
	arr := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	expected := []byte{'h', 'a', 'n', 'n', 'a', 'H'}
	reverseString(arr)
	fmt.Printf("nums, answer %v, %v \n", arr, expected)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_reverseString_1(t *testing.T) {
	// t.Skip()
	arr := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	reverseString(arr)

	fmt.Printf("nums, answer %v, %v \n", arr, expected)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}
