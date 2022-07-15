package maximum_subarray

func maxSubArray(nums []int) int {
	res, dp := nums[0], nums[0]

	for _, v := range nums {
		dp = getMax(v, v+dp)
		res = getMax(res, dp)
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
