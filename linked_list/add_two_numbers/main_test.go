package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	list1  []int
	list2  []int
	output []int
}{
	{
		list1:  []int{2, 4, 3},
		list2:  []int{5, 6, 4},
		output: []int{7, 0, 8},
	},
	{
		list1:  []int{0},
		list2:  []int{0},
		output: []int{0},
	},
	{
		list1:  []int{9, 9, 9, 9, 9, 9, 9},
		list2:  []int{9, 9, 9, 9},
		output: []int{8, 9, 9, 9, 0, 0, 0, 1},
	},
}

func Test_AddTwoList(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headList1 := convertArrToLinkedList(v.list1)
		headList2 := convertArrToLinkedList(v.list2)
		expectedOutput := convertArrToLinkedList(v.output)

		// traverse(headList1)
		// traverse(headList2)

		headOutput := addTwoNumbers(headList1, headList2)
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
