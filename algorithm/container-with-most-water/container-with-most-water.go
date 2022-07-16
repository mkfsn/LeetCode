package container_with_most_water

func maxArea(height []int) int {
	max := 0

	left, right := 0, len(height)-1

	for left < right {
		y1, y2 := height[left], height[right]
		x1, x2 := left, right

		max = getMax(max, (x2-x1)*getMin(y1, y2))
		if y1 < y2 {
			left++
		} else {
			right--
		}
	}

	return max
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 1 0 0 0 0 0 0 2 2 -> 8
// 1 1 2 2 0 0 0 0 0 0 0 3 3 -> 20
