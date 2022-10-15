package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	root   []interface{}
	output []int
}{
	{
		root:   []interface{}{1, nil, 2, 3},
		output: []int{1, 2, 3},
	},
	{
		root:   []interface{}{1},
		output: []int{1},
	},
	{
		root:   []interface{}{},
		output: []int{},
	},
}

func Test_PreorderTraversal(t *testing.T) {
	// reverseList
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		// convert array to node
		rootList := constructTreeNode(v.root, 0, len(v.root))
		expectedOutput := v.output

		// traverse(rootList)
		// traverse(expectedOutput)

		rootOutput := preorderTraversal(rootList)
		fmt.Print("hasil %v \n", rootOutput)

		fmt.Printf("isi output %v , %v > ", rootOutput, expectedOutput)
		assert.Equal(t, rootOutput, expectedOutput)
	}
}
