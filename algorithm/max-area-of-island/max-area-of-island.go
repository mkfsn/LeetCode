package max_area_of_island

func maxAreaOfIsland(grid [][]int) int {
	n, m := len(grid), len(grid[0])

	marks := make([][]int, n)
	for i := range grid {
		marks[i] = make([]int, m)
	}

	id := 0
	maxArea := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}

			if marks[i][j] != 0 { // travelled
				continue
			}

			id++

			area := discover(id, i, j, n, m, grid, marks)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func discover(id int, i int, j int, n int, m int, grid [][]int, marks [][]int) int {
	marks[i][j] = id

	area := 1

	// left
	if j-1 >= 0 && grid[i][j-1] == 1 && marks[i][j-1] == 0 {
		area += discover(id, i, j-1, n, m, grid, marks)
	}

	// right
	if j+1 < m && grid[i][j+1] == 1 && marks[i][j+1] == 0 {
		area += discover(id, i, j+1, n, m, grid, marks)
	}

	// up
	if i-1 >= 0 && grid[i-1][j] == 1 && marks[i-1][j] == 0 {
		area += discover(id, i-1, j, n, m, grid, marks)
	}

	// down
	if i+1 < n && grid[i+1][j] == 1 && marks[i+1][j] == 0 {
		area += discover(id, i+1, j, n, m, grid, marks)
	}

	return area
}
