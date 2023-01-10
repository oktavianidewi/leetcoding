package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	list1  []int
	output []int
}{
	{
		list1:  []int{1},
		output: []int{1},
	},
	{
		list1:  []int{},
		output: []int{},
	},
	{
		list1:  []int{1, 2},
		output: []int{2, 1},
	},
	{
		list1:  []int{1, 2, 3},
		output: []int{2, 1, 3},
	},
	{
		list1:  []int{1, 2, 3, 4},
		output: []int{2, 1, 4, 3},
	},
}

func TestSwapNodeInPairs(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to linked-list
		headList1 := convertArrToLinkedList(v.list1)
		// expectedOutput := convertArrToLinkedList(v.output)

		fmt.Print("before swap \n")
		traverse(headList1)

		swapPairs(headList1)

		fmt.Print("after swap \n")
		traverse(headList1)
		// fmt.Print("hasil ")
		// traverse(headOutput)

		// fmt.Printf("isi output %v , %v > ", headOutput, expectedOutput)
		// for headOutput != nil {
		// 	assert.Equal(t, headOutput, expectedOutput)
		// 	headOutput = headOutput.Next
		// 	expectedOutput = expectedOutput.Next
		// }
	}
}
