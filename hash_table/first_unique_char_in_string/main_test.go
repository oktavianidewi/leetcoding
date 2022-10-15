package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstUniqChar_0(t *testing.T) {
	// t.Skip()
	s := "leetcode"
	expected := 0
	answer := firstUniqChar(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_firstUniqChar_1(t *testing.T) {
	// t.Skip()
	s := "loveleetcode"
	expected := 2
	answer := firstUniqChar(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_firstUniqChar_2(t *testing.T) {
	// t.Skip()
	s := "aabb"
	expected := -1
	answer := firstUniqChar(s)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}
