package arrays

import (
	"testing"
)

func TestFindSecondLargest(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "should return second largest number from an unsorted array",
			input:    []int{1, 5, 3, 9, 7},
			expected: 7,
		},
		{
			name:     "should return smaller number when array has only two elements",
			input:    []int{10, 20},
			expected: 10,
		},
		{
			name:     "should return second largest when array is sorted in ascending order",
			input:    []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "should return second largest when array is sorted in descending order",
			input:    []int{9, 8, 7, 6, 5},
			expected: 8,
		},
		{
			name:     "should return second largest distinct number when duplicates are present",
			input:    []int{10, 10, 9, 8},
			expected: 9,
		},
		{
			name:     "should return -1 when all elements in array are equal",
			input:    []int{5, 5, 5, 5},
			expected: -1,
		},
		{
			name:     "should correctly return second largest number when array contains negative numbers",
			input:    []int{-10, -20, -30, -5},
			expected: -10,
		},
		{
			name:     "should return -1 when array contains less than 2 elements",
			input:    []int{10},
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			result := findSecondLargest(tc.input)

			if result != tc.expected {
				t.Errorf(
					"FindSecondLargest(%v) = %d, want %d",
					tc.input,
					result,
					tc.expected,
				)
			}
		})
	}
}
