package longest_palindromic_subsequence

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			k := n - j - 1

			if s[i] == s[j] {
				if i > 0 && k > 0 {
					dp[i][k] = 1 + dp[i-1][k-1]
				} else {
					dp[i][k] = 1
				}
			} else {

				x := 0
				if i > 0 {
					x = dp[i-1][k]
				}

				y := 0
				if k > 0 {
					y = dp[i][k-1]
				}

				dp[i][k] = getMax(x, y)
			}

		}
	}

	return dp[n-1][n-1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
