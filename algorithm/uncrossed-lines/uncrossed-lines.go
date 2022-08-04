package uncrossed_lines

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 + dp[i-1][j-1]
				}
			} else {
				x := 0
				if i > 0 {
					x = dp[i-1][j]
				}

				y := 0
				if j > 0 {
					y = dp[i][j-1]
				}

				dp[i][j] = getMax(x, y)
			}
		}
	}

	return dp[len(nums1)-1][len(nums2)-1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1 4 2
// |
// 1 2 4
// Ans: 2

//  2 5 1 2 5
//    |     |
// 10 5 2 1 5 2
// Ans: 3

// 1 2 3 4 5
// 2 3 4 5 6
// Ans: 4

// 1 3 7 1 7 5
// |
// 1 9 2 5 1
// Ans: 2
