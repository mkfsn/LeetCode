package kth_missing_positive_number

func findKthPositive(arr []int, k int) int {
	low, high := 0, len(arr)-1

	for low < high {
		i := (high + low) / 2
		v := arr[i]

		if v-i <= k {
			low = i + 1
		} else if v-i > k {
			high = i
		}
	}

	if arr[low]-low > k {
		return low + k
	}

	return low + k + 1
}
