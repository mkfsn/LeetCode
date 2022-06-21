package longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[j]+1, dp[i])
			}
		}
	}

	res := 1

	for _, v := range dp {
		res = max(res, v)
	}

	return res
}

func max(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
