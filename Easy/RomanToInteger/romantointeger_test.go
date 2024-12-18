package romantointeger

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
			name:     "III",
			input:    "III",
			expected: 3,
		},
		{
			name:     "LVIII",
			input:    "LVIII",
			expected: 58,
		},
		{
			name:     "MCMXCIV",
			input:    "MCMXCIV",
			expected: 1994,
		},
	}
)

func TestRomanToIntFirst(t *testing.T) {

	for _, tc := range testCases {
		assert := assert.New(t)

		t.Run(tc.name, func(t *testing.T) {
			out := RomanToIntFirst(tc.input)
			assert.Equal(tc.expected, out, "expect %d to be %d but got %d", out, tc.expected, out)
		})
	}

}
