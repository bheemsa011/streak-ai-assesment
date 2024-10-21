package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindPairsInArray(t *testing.T) {

	testCases := []struct {
		numbers  []int
		target   int
		expected [][]int
	}{
		{
			numbers:  []int{1, 2, 3, 4, 5},
			target:   6,
			expected: [][]int{{1, 3}, {0, 4}},
		},
		{numbers: []int{1, 2, 3},
			target:   10,
			expected: [][]int{},
		},
		{numbers: []int{},
			target:   0,
			expected: [][]int{},
		},
	}

	for _, test := range testCases {

		result := FindPairsInArray(test.numbers, test.target)
		assert.Equal(t, test.expected, result)

	}
}
