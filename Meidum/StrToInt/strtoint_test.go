package strtoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "spaces only",
			input:    "   ",
			expected: 0,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "explicit positive",
			input:    "+12",
			expected: 12,
		},
		{
			name:     "multiple signs",
			input:    "-+12",
			expected: 0,
		},
		{
			name:     "string with non-numerical characters",
			input:    "1337c0d3",
			expected: 1337,
		},
		{
			name:     "random string -1",
			input:    "words and 987",
			expected: 0,
		},
		{
			name:     "random string -2",
			input:    "0-1",
			expected: 0,
		},
		{
			name:     "negative out of bound",
			input:    "-91283472332",
			expected: -2147483648,
		},
		{
			name:     "positive out of bound",
			input:    "91283472332",
			expected: 2147483647,
		},
		{
			name:     "positive int",
			input:    "12",
			expected: 12,
		},
		{
			name:     "negative int",
			input:    "-12",
			expected: -12,
		},
	}
)

func TestMyAtoi(t *testing.T) {

	for _, tc := range testCases {
		assert := assert.New(t)
		t.Run(tc.name, func(t *testing.T) {
			out := MyAtoi(tc.input)
			assert.Equal(tc.expected, out, "expect %d to be %d but got %d", out, tc.expected, out)
		})

	}

}
