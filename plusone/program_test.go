package plusone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CleanTrailingZero(t *testing.T) {
	// var input digits
	// input = []int{1, 2, 3}
	// input := []int{1, 2, 3}
	input := []int{0, 0, 0, 9}
	expected := []int{9}

	output := cleanTrailingZero(input)
	assert.Equal(t, expected, output)
	// for k, v := range output {
	// 	assert.Equal(t, v, input[k])
	// }
}

func Test_PlusOne(t *testing.T) {
	// var input digits
	// input = []int{1, 2, 3}
	input := []int{1, 2, 3}

	output := plusone(input)
	assert.Equal(t, output, input)
	// for k, v := range output {
	// 	assert.Equal(t, v, input[k])
	// }
}
