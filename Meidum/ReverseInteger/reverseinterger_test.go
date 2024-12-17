package reverseinteger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "positive int ending with 0",
			input:    120,
			expected: 21,
		},
		{
			name:     "positive int",
			input:    121,
			expected: 121,
		},
		{
			name:     "negative int",
			input:    -123,
			expected: -321,
		},
		{
			name:     "positive out of bound after reverse",
			input:    1534236469,
			expected: 0,
		},
		{
			name:     "negative out of bound after reverse",
			input:    -1534236469,
			expected: 0,
		},
	}
)

func TestReverse(t *testing.T) {
	for _, tc := range testCases {
		assert := assert.New(t)
		t.Run(tc.name, func(t *testing.T) {
			out := Reverse(tc.input)
			assert.Equal(tc.expected, out, "expect %d to be %d, but got %d", out, tc.expected, out)
		})
	}
}
