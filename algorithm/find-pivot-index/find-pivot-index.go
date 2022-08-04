package find_pivot_index

func pivotIndex(nums []int) int {
	sums := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		x := 0
		if i < len(nums)-1 {
			x = sums[i+1]
		}
		sums[i] = x + nums[i]
	}

	accumulate := 0
	for i := 0; i < len(nums)-1; i++ {
		if sums[i]-nums[i] == accumulate {
			return i
		}

		accumulate += nums[i]
	}

	return -1
}
