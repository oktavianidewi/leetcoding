package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ModifyString(t *testing.T) {
	str := "de75s1rev2er"
	answer := ModifyString(str)
	expected := "reversed"
	fmt.Printf("nums, answer %v, %v \n", answer, expected)
	assert.Equal(t, expected, answer)

}
