package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords_0(t *testing.T) {
	// t.Skip()
	s := "the sky is blue"
	expected := "blue is sky the"
	answer := reverseWords(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)

	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_reverseWords_1(t *testing.T) {
	// t.Skip()
	s := "  hello world  "
	expected := "world hello"
	answer := reverseWords(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}
