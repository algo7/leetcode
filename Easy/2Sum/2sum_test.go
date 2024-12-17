package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testCases = []struct {
		name     string
		target   int
		input    []int
		expected []int
	}{
		{
			name:     "test - 4 elements",
			target:   9,
			input:    []int{2, 7, 11, 15},
			expected: []int{0, 1},
		},
		{
			name:     "test - 3 elements",
			target:   6,
			input:    []int{3, 2, 4},
			expected: []int{1, 2},
		},
		{
			name:     "test - duplicate elements",
			target:   6,
			input:    []int{3, 3},
			expected: []int{0, 1},
		},
	}
)

func TestTwoSumDoubleLoop(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			assert := assert.New(t)

			r := TwoSumDoubleLoop(tc.input, tc.target)

			assert.ElementsMatch(r, tc.expected, "expect %v to be %v, but got %v", r, tc.expected, r)

		})
	}

}

func TestTwoSumMap(t *testing.T) {

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			assert := assert.New(t)

			r := TwoSumMap(tc.input, tc.target)

			assert.ElementsMatch(r, tc.expected, "expect %v to be %v, but got %v", r, tc.expected, r)

		})
	}

}
