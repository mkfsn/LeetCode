package sort_array_by_parity_ii

func sortArrayByParityII(nums []int) []int {
	i, j, k := 0, 1, len(nums)-1

	for i <= k && j <= k {
		x := nums[k]

		// fmt.Printf("(i=%d,j=%d,k=%d,x=%d): %v\n", i, j, k, x, nums)

		if x%2 == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			i += 2
		} else { // x % 2 == 1
			nums[j], nums[k] = nums[k], nums[j]
			j += 2
		}
	}

	return nums
}
