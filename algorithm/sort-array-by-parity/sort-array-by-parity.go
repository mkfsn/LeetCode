package sort_array_by_parity

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for j > i {
		if nums[j]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		} else {
			j--
		}
	}
	return nums
}
