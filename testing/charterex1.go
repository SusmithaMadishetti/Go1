package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumIt(t *testing.T) {
	assert.Equal(t, 12, SumIt(3, 5))
	assert.Equal(t, 22, SumIt(4, 7))
	assert.Equal(t, 42, SumIt(3, 9))
}

// Returns an integer that is the sum of the integers between a and b.
//
// - If 3 and 5 are passed in the result should be 12: (3 + 4 + 5) = 12.
// - If 4 and 7 are passed in the result should be 22: (4 + 5 + 6 + 7) = 22.
//
// Complete the function below to make the test cases pass.
//
func SumIt(a, b int) int {
	c := 0
	if a > b {
		a, b = b, a
	}
	for i := a; i <= b; i++ {
		c = c + i
	}
	return c
}
