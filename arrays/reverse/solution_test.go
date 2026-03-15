package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "should reverse array with odd number of elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "should reverse array with even number of elements",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "should return same array when array has single element",
			input:    []int{10},
			expected: []int{10},
		},
		{
			name:     "should return empty array when input is empty",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "should reverse array containing negative numbers",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-4, -3, -2, -1},
		},
		{
			name:     "should reverse array with mixed positive and negative numbers",
			input:    []int{10, -5, 3, 0},
			expected: []int{0, 3, -5, 10},
		},
		{
			name:     "should reverse already reversed array",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			input := append([]int{}, tc.input...)
			Reverse(input)

			assert.Equal(t, tc.expected, input)
		})
	}
}
