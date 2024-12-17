package fizzbuzz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = []struct {
		name     string
		input    int
		expected []string
	}{
		{
			name:     "test1",
			input:    3,
			expected: []string{"1", "2", "Fizz"},
		},
		{
			name:     "test2",
			input:    5,
			expected: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:     "test3",
			input:    15,
			expected: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
)

func TestFizzBuzzFirst(t *testing.T) {

	for _, tc := range testCases {

		assert := assert.New(t)

		out := FizzBuzzFirst(tc.input)

		t.Run(tc.name, func(t *testing.T) {

			assert.ElementsMatch(out, tc.expected, "expect %v to be %v, but got %v", out, tc.expected, out)

		})

	}

}

func TestFizzBuzzSecond(t *testing.T) {

	for _, tc := range testCases {

		assert := assert.New(t)

		t.Run(tc.name, func(t *testing.T) {

			out := FizzBuzzSecond(tc.input)

			assert.ElementsMatch(out, tc.expected, "expect %v to be %v, but got %v", out, tc.expected, out)

		})

	}

}
