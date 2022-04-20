package binary_search

func search(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		i := (low + high) / 2
		v := nums[i]

		if v == target {
			return i
		} else if v < target {
			low = i + 1
		} else {
			high = i - 1
		}
	}

	return -1
}
