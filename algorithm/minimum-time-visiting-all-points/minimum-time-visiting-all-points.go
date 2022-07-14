package minimum_time_visiting_all_points

func minTimeToVisitAllPoints(points [][]int) int {
	res := 0
	for i := range points {
		if i == len(points)-1 {
			break
		}

		p1, p2 := points[i], points[i+1]

		res += getMax(
			abs(p2[0]-p1[0]),
			abs(p2[1]-p1[1]),
		)
	}

	return res
}

func abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
