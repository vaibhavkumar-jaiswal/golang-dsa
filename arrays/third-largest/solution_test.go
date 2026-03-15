package arrays

import (
	"math"
	"testing"
)

func TestFindThirdLargest(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "should return third largest element from unsorted array",
			input:    []int{1, 14, 2, 16, 10, 20},
			expected: 14,
		},
		{
			name:     "should return third largest when array contains negative numbers",
			input:    []int{19, -10, 20, 14, 2, 16, 10},
			expected: 16,
		},
		{
			name:     "should return third largest when array is sorted in ascending order",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: 4,
		},
		{
			name:     "should return third largest when array is sorted in descending order",
			input:    []int{9, 8, 7, 6, 5, 4},
			expected: 7,
		},
		{
			name:     "should return smallest element when array contains exactly three elements",
			input:    []int{5, 10, 15},
			expected: 5,
		},
		{
			name:     "should handle array with duplicate values",
			input:    []int{10, 20, 20, 30, 40},
			expected: 20,
		},
		{
			name:     "should return -1 when array has fewer than three elements",
			input:    []int{10, 20},
			expected: -1,
		},
		{
			name:     "should return -1 when array has exactly two elements",
			input:    []int{5, 7},
			expected: -1,
		},
		{
			name:     "should return -1 when all elements are same",
			input:    []int{10, 10, 10, 10},
			expected: -1,
		},
		{
			name:     "should return -1 when only two distinct elements exist",
			input:    []int{10, 20, 20, 20},
			expected: -1,
		},
		{
			name:     "should correctly return third largest when numbers include math.MinInt",
			input:    []int{math.MinInt, -5, -10, -1},
			expected: -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := findThirdLargest(tt.input)

			if result != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, result)
			}
		})
	}
}
