package find_the_highest_altitude

func largestAltitude(gain []int) int {
	max, cur := 0, 0

	for _, v := range gain {
		cur += v
		if cur > max {
			max = cur
		}
	}

	return max
}
