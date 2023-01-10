package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  []int
	output []int
}{
	{
		input:  []int{1, 2, 3, 4, 5},
		output: []int{1, 3, 5, 2, 4},
	},
	{
		input:  []int{2, 1, 3, 5, 6, 4, 7},
		output: []int{2, 3, 6, 7, 1, 5, 4},
	},
	{
		input:  []int{1},
		output: []int{1},
	},
	{
		input:  []int{},
		output: []int{},
	},
}

func Test_OddEvenList(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headInput := convertArrToLinkedList(v.input)
		expectedOutput := convertArrToLinkedList(v.output)

		headOutput := oddEvenList(headInput)
		fmt.Printf("isi output %v , %v > ", headOutput, expectedOutput)

		// traverse(headInput)

		// traverse(headOutput)

		for headOutput != nil {
			assert.Equal(t, headOutput, expectedOutput)
			headOutput = headOutput.Next
			expectedOutput = expectedOutput.Next
		}
	}
}
