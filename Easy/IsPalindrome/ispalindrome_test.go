package ispalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "normal int",
			input:    121,
			expected: true,
		},
		{
			name:     "negative int",
			input:    -121,
			expected: false,
		},
		{
			name:     "int ending with 0",
			input:    10,
			expected: false,
		},
	}
)

func TestIsPalindromeString(t *testing.T) {

	for _, tc := range testCases {
		assert := assert.New(t)
		t.Run(tc.name, func(t *testing.T) {

			out := IsPalindromeString(tc.input)
			assert.Equal(tc.expected, out, "expect %t to be %t but got %t", out, tc.expected, out)

		})
	}

}

func TestIsPalindromeStringShort(t *testing.T) {

	for _, tc := range testCases {
		assert := assert.New(t)
		t.Run(tc.name, func(t *testing.T) {

			out := IsPalindromeStringShort(tc.input)
			assert.Equal(tc.expected, out, "expect %t to be %t but got %t", out, tc.expected, out)

		})
	}

}

func TestIsPalindromeMod(t *testing.T) {

	for _, tc := range testCases {
		assert := assert.New(t)
		t.Run(tc.name, func(t *testing.T) {

			out := IsPalindromeMod(tc.input)
			assert.Equal(tc.expected, out, "expect %t to be %t but got %t", out, tc.expected, out)

		})
	}

}
