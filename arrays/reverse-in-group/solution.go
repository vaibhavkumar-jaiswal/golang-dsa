package arrays

func reverseInGroups(arr []int, k int) []int {
	if k <= 1 {
		return arr
	}

	n := len(arr)

	for start := 0; start < n; start += k {

		end := start + k - 1
		if end >= n {
			end = n - 1
		}

		for i, j := start, end; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	return arr
}
