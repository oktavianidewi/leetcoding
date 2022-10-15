package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_moveZeroes_0(t *testing.T) {
	// t.Skip()
	nums := []int{0, 1, 0, 3, 12}
	expected := "eht yks si eulb"
	answer := moveZeroes(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)

	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_moveZeroes_1(t *testing.T) {
	// t.Skip()
	s := "  hello world  "
	expected := "olleh dlrow"
	answer := moveZeroes(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_moveZeroes_2(t *testing.T) {
	// t.Skip()
	s := "Let's take LeetCode contest"
	expected := "s'teL ekat edoCteeL tsetnoc"
	answer := moveZeroes(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}

func Test_moveZeroes_3(t *testing.T) {
	// t.Skip()
	s := "God Ding"
	expected := "doG gniD"
	answer := moveZeroes(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for i, _ := range expected {
	// 	assert.Equal(t, expected[i], answer[i])
	// }
}
