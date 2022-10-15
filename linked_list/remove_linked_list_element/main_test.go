package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input  []int
	n      int
	output []int
}{
	{
		input:  []int{1, 2, 6, 3, 4, 5, 6},
		n:      6,
		output: []int{1, 2, 3, 4, 5},
	},
	{
		input:  []int{6, 2, 6, 3, 4, 5, 6},
		n:      6,
		output: []int{2, 3, 4, 5},
	},
	{
		input:  []int{7, 7, 7, 7},
		n:      7,
		output: []int{},
	},
	{
		input:  []int{},
		n:      1,
		output: []int{},
	},
}

func Test_RemoveElementsFromLinkedList(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headInput := convertArrToLinkedList(v.input)
		expectedOutput := convertArrToLinkedList(v.output)

		headOutput := removeElements(headInput, v.n)
		fmt.Printf("isi output %v , %v > ", headOutput, expectedOutput)

		traverse(headOutput)

		for headOutput != nil {
			assert.Equal(t, headOutput, expectedOutput)
			headOutput = headOutput.Next
			expectedOutput = expectedOutput.Next
		}
	}
}
