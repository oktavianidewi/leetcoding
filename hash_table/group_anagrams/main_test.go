package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_groupAnagrams_0(t *testing.T) {
	// t.Skip()
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"bat"},
		{"nat", "tan"},
		{"ate", "eat", "tea"},
	}
	answer := groupAnagrams(strs)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for _, v := range expected {
	// 	assert.Equal(t, v, answer)
	// }
}

func Test_groupAnagrams_1(t *testing.T) {
	// t.Skip()
	strs := []string{""}
	expected := [][]string{
		{""},
	}
	answer := groupAnagrams(strs)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for _, v := range expected {
	// 	assert.Equal(t, v, answer)
	// }
}

func Test_groupAnagrams_2(t *testing.T) {
	// t.Skip()
	strs := []string{"a"}
	expected := [][]string{
		{"a"},
	}
	answer := groupAnagrams(strs)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
	// for _, v := range expected {
	// 	assert.Equal(t, v, answer)
	// }
}
