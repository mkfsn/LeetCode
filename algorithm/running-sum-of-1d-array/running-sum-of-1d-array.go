package running_sum_of_1d_array

func runningSum(nums []int) []int {
	sum := 0
	for i, num := range nums {
		nums[i] = sum + num
		sum = nums[i]
	}
	return nums
}
