package arrays

func RotateLeftByD(arr []int, d int) []int {
	n := len(arr)

	if n == 0 {
		return arr
	}

	d = d % n
	if d == 0 {
		return arr
	}

	left := arr[:d]

	arr = append(arr[d:], left...)

	return arr
}

func optimalSolution(arr []int, d int) {
	n := len(arr)

	if n == 0 {
		return
	}

	d = d % n

	reverse(arr, 0, d-1)
	reverse(arr, d, n-1)
	reverse(arr, 0, n-1)
}

func reverse(arr []int, left, right int) {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
}
