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
		output: []int{5, 4, 3, 2, 1},
	},
	{
		input:  []int{1, 2},
		output: []int{2, 1},
	},
	{
		input:  []int{},
		output: []int{},
	},
}

func Test_ReverseLinkedList(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headInput := convertArrToLinkedList(v.input)
		expectedOutput := convertArrToLinkedList(v.output)

		headOutput := reverseList(headInput)
		fmt.Printf("isi output %v , %v > ", headOutput, expectedOutput)

		for headOutput != nil {
			assert.Equal(t, headOutput, expectedOutput)
			headOutput = headOutput.Next
			expectedOutput = expectedOutput.Next
		}
	}
}
