package unique_paths

func uniquePaths(m int, n int) int {
	area := make([][]int, m)

	for i := range area {
		area[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				area[i][j] = 1
			} else {
				area[i][j] = area[i-1][j] + area[i][j-1]
			}
		}
	}

	return area[m-1][n-1]
}

// 1 1 1 1  1  1  1
// 1 2 3 4  5  6  7
// 1 3 6 10 15 21 28
