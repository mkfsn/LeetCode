package find_the_middle_index_in_array

func findMiddleIndex(nums []int) int {
	sum := make([]int, len(nums))

	n := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		sum[i] = n
		n += nums[i]
	}

	n = 0
	for i := 0; i < len(nums); i++ {
		if sum[i] == n {
			return i
		}
		n += nums[i]
	}

	return -1
}

// 11:27
// 11:38
