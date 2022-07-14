package number_of_islands

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	marks := make([][]int, len(grid))
	for i, row := range grid {
		marks[i] = make([]int, len(row))
	}

	cnt := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				continue
			}

			if marks[i][j] != 0 {
				continue // has travelled
			}

			cnt++
			markIsland(i, j, cnt, marks, grid)

			printMarks(marks)
		}
	}

	return cnt
}

func printMarks(marks [][]int) {
	for _, row := range marks {
		for _, v := range row {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
	fmt.Println()
}

func markIsland(i int, j int, cnt int, marks [][]int, grid [][]byte) {
	marks[i][j] = cnt

	n, m := len(grid), len(grid[0])

	// up
	if i-1 >= 0 && marks[i-1][j] == 0 && grid[i-1][j] == '1' {
		markIsland(i-1, j, cnt, marks, grid)
	}

	// down
	if i+1 < n && marks[i+1][j] == 0 && grid[i+1][j] == '1' {
		markIsland(i+1, j, cnt, marks, grid)
	}

	// left
	if j-1 >= 0 && marks[i][j-1] == 0 && grid[i][j-1] == '1' {
		markIsland(i, j-1, cnt, marks, grid)
	}

	// right
	if j+1 < m && marks[i][j+1] == 0 && grid[i][j+1] == '1' {
		markIsland(i, j+1, cnt, marks, grid)
	}
}
