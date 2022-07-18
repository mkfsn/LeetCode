package flood_fill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	newImage := make([][]int, len(image))

	n, m := len(image), len(image[0])

	for i, row := range image {
		newImage[i] = make([]int, len(row))
		for j, c := range row {
			newImage[i][j] = c
		}
	}

	flood(image, sr, sc, n, m, color, newImage, image[sr][sc])

	return newImage
}

func flood(oldImage [][]int, x int, y int, n int, m int, color int, newImage [][]int, target int) {
	newImage[x][y] = color

	if x-1 >= 0 && oldImage[x-1][y] == target && newImage[x-1][y] != color {
		flood(oldImage, x-1, y, n, m, color, newImage, target)
	}

	if y-1 >= 0 && oldImage[x][y-1] == target && newImage[x][y-1] != color {
		flood(oldImage, x, y-1, n, m, color, newImage, target)
	}

	if x+1 < n && oldImage[x+1][y] == target && newImage[x+1][y] != color {
		flood(oldImage, x+1, y, n, m, color, newImage, target)
	}

	if y+1 < m && oldImage[x][y+1] == target && newImage[x][y+1] != color {
		flood(oldImage, x, y+1, n, m, color, newImage, target)
	}
}

// 13:44
// 13:49 toilet
// 13:55 back
// 14:01
