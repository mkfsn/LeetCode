package count_negative_numbers_in_a_sorted_matrix

func countNegatives(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	i, j := m-1, 0
	cnt := 0

	for i >= 0 && j < n {
		if grid[i][j] < 0 {
			cnt += n - j
			i--
		} else {
			j++
		}
	}

	return cnt
}

//  4  3  2 -1
//  3  2  1 -1
//  1  1 -1 -1
// -1 -1 -2 -3
