package can_make_arithmetic_progression_from_sequence

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) <= 2 {
		return true
	}

	min, max := arr[0], arr[0]

	for i := 1; i < len(arr); i++ {
		min = getMin(min, arr[i])
		max = getMax(max, arr[i])
	}

	if (max-min)%(len(arr)-1) != 0 {
		return false
	}

	counts := make(map[int]bool)
	diff := (max - min) / (len(arr) - 1)

	for _, v := range arr {
		counts[v] = true
	}

	for i := 1; i < len(arr); i++ {
		if _, ok := counts[min+diff*i]; !ok {
			return false
		}
	}

	return true
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a

	}
	return b
}
