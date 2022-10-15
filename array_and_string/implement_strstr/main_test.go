package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strStr_1(t *testing.T) {
	t.Skip()
	haystack := "hello"
	needle := "ll"
	expected := 2
	answer := strStr(haystack, needle)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_strStr_2(t *testing.T) {
	t.Skip()
	haystack := "aaaaa"
	needle := "bba"
	expected := -1
	answer := strStr(haystack, needle)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_strStr_3(t *testing.T) {
	t.Skip()
	haystack := "aaa"
	needle := "aaaa"
	expected := -1
	answer := strStr(haystack, needle)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_strStr_4(t *testing.T) {
	t.Skip()
	haystack := "mississippi"
	needle := "issip"
	expected := 4
	answer := strStr(haystack, needle)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_strStr_5(t *testing.T) {
	// t.Skip()
	haystack := "mississippi"
	needle := "issipi"
	expected := -1
	answer := strStr(haystack, needle)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}
