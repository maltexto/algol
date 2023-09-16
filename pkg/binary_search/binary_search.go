package binary_search

func Run(values []int, target int) int {
	low := 0
	high := len(values) - 1

	for low <= high {
		mid := (low + high) / 2
		if values[mid] == target {
			return mid
		} else if values[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
