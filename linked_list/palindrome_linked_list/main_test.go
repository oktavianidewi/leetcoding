package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  []int
	output bool
}{
	{
		input:  []int{1, 2, 3, 4, 5},
		output: false,
	},
	{
		input:  []int{2, 1, 3, 5, 6, 4, 7},
		output: false,
	},
	{
		input:  []int{5, 5, 5, 5},
		output: true,
	},
	{
		input:  []int{1, 2, 2, 1},
		output: true,
	},
}

func Test_OddEvenList(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headInput := convertArrToLinkedList(v.input)
		expectedOutput := v.output

		headOutput := isPalindrome(headInput)
		fmt.Printf("isi output %v , %v > ", headOutput, expectedOutput)
		assert.Equal(t, headOutput, expectedOutput)
	}
}
