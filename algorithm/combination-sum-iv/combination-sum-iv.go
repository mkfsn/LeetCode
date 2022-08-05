package combination_sum_iv

func combinationSum4(nums []int, target int) int {
	history := make(map[int]int)
	return dfs(nums, target, history)
}

func dfs(nums []int, target int, history map[int]int) int {
	if target <= 0 {
		return 0
	}

	if v, ok := history[target]; ok {
		return v
	}

	cnt := 0
	for _, n := range nums {
		if n == target {
			cnt += 1
			continue
		}

		cnt += dfs(nums, target-n, history)
	}

	history[target] = cnt

	return cnt
}

// nums = [1,2,3], target = 4
// ->
//    1 + combinationSum4([1,2,3], 3)
//      -> 1 + combinationSum4([1,2,3], 2)
//           -> 1 + ombinationSum4([1,2,3], 1)
//                -> 1
//           -> 2
//      -> 2 + combinationSum4([1,2,3], 1)
//           -> 1
//      -> 3
//    2 + combinationSum4([1,2,3], 2)
//      -> 1 + combinationSum4([1,2,3], 1)
//           -> 1
//      -> 2
//    3 + combinationSum4([1,2,3], 1)
//      -> 1
// dp[i] = (dp[i-1] + dp[i-2] + ... + dp[1] + dp[0])
// dp[0] = 1
// dp[1] = dp[0] = 1
// dp[2] = dp[1] + dp[0] = 2
// dp[3] = dp[0] + dp[1] + dp[2] = 4

// nums: [2, 3, 4], target = 7
// dp[0] = 1
// dp[1] = dp[0] = 0
// dp[2] = dp[0] = 1
// dp[3] = dp[0] = 1
// dp[4] = dp[0] = 1
// dp[5] = dp[2] + dp[3] = 2
// dp[6] = 2+4, 3+3 = 2
