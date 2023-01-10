package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	head   []int
	k      int
	output []int
}{
	{
		head:   []int{1, 2, 3, 4, 5},
		k:      2,
		output: []int{4, 5, 1, 2, 3},
	},
	{
		head:   []int{0, 1, 2},
		k:      4,
		output: []int{2, 0, 1},
	},
}

func Test_RotateList(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headList := convertArrToLinkedList(v.head)
		expectedOutput := convertArrToLinkedList(v.output)

		// traverse(headList)
		// traverse(expectedOutput)

		headOutput := rotateRight(headList, v.k)
		fmt.Print("hasil ")
		traverse(headOutput)

		fmt.Printf("isi output %v , %v > ", headOutput, expectedOutput)
		for headOutput != nil {
			assert.Equal(t, headOutput, expectedOutput)
			headOutput = headOutput.Next
			expectedOutput = expectedOutput.Next
		}
	}
}
