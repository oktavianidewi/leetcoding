package opec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader_Calculate(t *testing.T) {
	calculate := Calculate(10, 20)
	assert.Equal(t, calculate.result, 30)
}
