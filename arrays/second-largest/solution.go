package arrays

import "math"

func findSecondLargest(arr []int) int {
	if len(arr) < 2 {
		return -1
	}

	first := math.MinInt
	second := math.MinInt

	for _, value := range arr {
		if value > first {
			second = first
			first = value
		} else if value > second && value != first {
			second = value
		}
	}

	if second == math.MinInt {
		return -1
	}

	return second
}
