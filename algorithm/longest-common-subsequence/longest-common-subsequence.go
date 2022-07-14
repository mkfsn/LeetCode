package longest_common_subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	n, m := len(text1), len(text2)

	records := make([][]int, n)
	for i := range records {
		records[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if text1[i] == text2[j] {
				if i == 0 || j == 0 {
					records[i][j] = 1
				} else {
					records[i][j] = 1 + records[i-1][j-1]
				}
			} else {
				x := 0
				if i > 0 {
					x = records[i-1][j]
				}

				y := 0
				if j > 0 {
					y = records[i][j-1]
				}

				records[i][j] = getMax(x, y) // [i-1][j] || [i][j-1]
			}
		}
	}

	return records[n-1][m-1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
