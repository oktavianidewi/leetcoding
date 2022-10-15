package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	n      int
	output []int
}{
	{
		n:      0,
		output: []int{1},
	},
	{
		n:      1,
		output: []int{1, 1},
	},
	{
		n:      2,
		output: []int{1, 2, 1},
	},
	{
		n:      3,
		output: []int{1, 3, 3, 1},
	},
	{
		n:      4,
		output: []int{1, 4, 6, 4, 1},
	},
	{
		n:      5,
		output: []int{1, 5, 10, 10, 5, 1},
	},
}

func TestPascalTriangle(t *testing.T) {
	fmt.Printf("test cases %v\n", testCases)
	for _, v := range testCases {
		outputResult := getRow(v.n)
		fmt.Printf("isi output %v , expected %v > \n", outputResult, v.output)
		assert.Equal(t, outputResult, v.output)

		// for headOutput != nil {
		// 	assert.Equal(t, headOutput, expectedOutput)
		// 	headOutput = headOutput.Next
		// 	expectedOutput = expectedOutput.Next
		// }
	}
}
