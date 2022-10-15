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
		list1:  []int{1, 2, 4},
		list2:  []int{1, 3, 4},
		output: []int{1, 1, 2, 3, 4, 4},
	},
	{
		list1:  []int{},
		list2:  []int{},
		output: []int{},
	},
	{
		list1:  []int{},
		list2:  []int{5},
		output: []int{5},
	},
	{
		list1:  []int{1, 2, 3},
		list2:  []int{4, 5, 6},
		output: []int{1, 2, 3, 4, 5, 6},
	},
	{
		list1:  []int{1, 2, 4},
		list2:  []int{1, 3, 4},
		output: []int{1, 1, 2, 3, 4, 4},
	},
}

func Test_MergeTwoList(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headList1 := convertArrToLinkedList(v.list1)
		headList2 := convertArrToLinkedList(v.list2)
		expectedOutput := convertArrToLinkedList(v.output)

		// traverse(headList1)
		// traverse(headList2)

		headOutput := mergeTwoLists(headList1, headList2)
		// traverse(headOutput)

		fmt.Printf("isi output %v , %v > ", headOutput, expectedOutput)
		for headOutput != nil {
			assert.Equal(t, headOutput, expectedOutput)
			headOutput = headOutput.Next
			expectedOutput = expectedOutput.Next
		}
	}
}
