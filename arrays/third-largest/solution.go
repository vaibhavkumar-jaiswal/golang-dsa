package arrays

import "math"

func findThirdLargest(arr []int) int {
	if len(arr) < 3 {
		return -1
	}

	first, second, third := math.MinInt, math.MinInt, math.MinInt

	for _, value := range arr {
		if value > first {
			third = second
			second = first
			first = value
		} else if value > second && value != first {
			third = second
			second = value
		} else if value > third && value != first && value != second {
			third = value
		}
	}

	if third == math.MinInt {
		return -1
	}

	return third
}
