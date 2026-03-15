package arrays

import (
	"reflect"
	"testing"
)

func TestReverseInGroups(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "should reverse elements in groups of k when array size is multiple of k",
			input:    []int{1, 2, 3, 4, 5, 6},
			k:        3,
			expected: []int{3, 2, 1, 6, 5, 4},
		},
		{
			name:     "should reverse last group when remaining elements are less than k",
			input:    []int{1, 2, 3, 4, 5},
			k:        3,
			expected: []int{3, 2, 1, 5, 4},
		},
		{
			name:     "should reverse entire array when k is greater than array length",
			input:    []int{5, 6, 8, 9},
			k:        5,
			expected: []int{9, 8, 6, 5},
		},
		{
			name:     "should not modify array when k is one",
			input:    []int{1, 2, 3},
			k:        1,
			expected: []int{1, 2, 3},
		},
		{
			name:     "should return same array when array contains single element",
			input:    []int{7},
			k:        3,
			expected: []int{7},
		},
		{
			name:     "should return empty array when input array is empty",
			input:    []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "should return same array when k is zero",
			input:    []int{1, 2, 3},
			k:        0,
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			input := make([]int, len(tt.input))
			copy(input, tt.input)

			result := reverseInGroups(input, tt.k)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, input)
			}
		})
	}
}
