package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateLeftByD(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		d        int
		expected []int
	}{
		{
			name:     "should rotate array left by d positions when d is less than array length",
			input:    []int{1, 2, 3, 4, 5, 6},
			d:        2,
			expected: []int{3, 4, 5, 6, 1, 2},
		},
		{
			name:     "should rotate array correctly when d is greater than array length",
			input:    []int{1, 2, 3},
			d:        4,
			expected: []int{2, 3, 1},
		},
		{
			name:     "should return same array when d is zero",
			input:    []int{1, 2, 3},
			d:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "should return same array when d equals array length",
			input:    []int{1, 2, 3},
			d:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "should handle single element array when rotated",
			input:    []int{5},
			d:        10,
			expected: []int{5},
		},
		{
			name:     "should return empty array when input array is empty",
			input:    []int{},
			d:        3,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			input := make([]int, len(tt.input))
			copy(input, tt.input)

			result := RotateLeftByD(input, tt.d)

			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOptimalSolution(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		d        int
		expected []int
	}{
		{
			name:     "should rotate array left by d positions when d is less than array length",
			input:    []int{1, 2, 3, 4, 5, 6},
			d:        2,
			expected: []int{3, 4, 5, 6, 1, 2},
		},
		{
			name:     "should rotate array correctly when d is greater than array length",
			input:    []int{1, 2, 3},
			d:        4,
			expected: []int{2, 3, 1},
		},
		{
			name:     "should return same array when d is zero",
			input:    []int{1, 2, 3},
			d:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "should return same array when d equals array length",
			input:    []int{1, 2, 3},
			d:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "should handle single element array when rotated",
			input:    []int{5},
			d:        10,
			expected: []int{5},
		},
		{
			name:     "should return empty array when input array is empty",
			input:    []int{},
			d:        3,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			input := make([]int, len(tt.input))
			copy(input, tt.input)

			optimalSolution(input, tt.d)

			assert.Equal(t, tt.expected, input)
		})
	}
}
