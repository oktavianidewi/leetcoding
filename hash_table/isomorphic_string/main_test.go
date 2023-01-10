package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isIsomorphic_0(t *testing.T) {
	// t.Skip()
	s1 := "egg"
	s2 := "add"
	expected := true
	answer := isIsomorphic(s1, s2)
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_isIsomorphic_1(t *testing.T) {
	// t.Skip()
	s1 := "foo"
	s2 := "bar"
	expected := false
	answer := isIsomorphic(s1, s2)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_isIsomorphic_2(t *testing.T) {
	// t.Skip()
	s1 := "paper"
	s2 := "title"
	expected := true
	answer := isIsomorphic(s1, s2)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}

func Test_isIsomorphic_3(t *testing.T) {
	// t.Skip()
	s1 := "badc"
	s2 := "baba"
	expected := false
	answer := isIsomorphic(s1, s2)

	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)
}
