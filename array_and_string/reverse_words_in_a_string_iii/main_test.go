package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseWords_0(t *testing.T) {
	// t.Skip()
	s := "the sky is blue"
	expected := "eht yks si eulb"
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
	expected := "olleh dlrow"
	answer := reverseWords(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_reverseWords_2(t *testing.T) {
	// t.Skip()
	s := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"
	answer := reverseWords(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_reverseWords_3(t *testing.T) {
	// t.Skip()
	s := "God Ding"
	expected := "doG gniD"
	answer := reverseWords(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}
