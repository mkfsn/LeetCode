package check_whether_two_strings_are_almost_equivalent

func checkAlmostEquivalent(word1 string, word2 string) bool {
	frequencies := make(map[int32]int)

	for _, c := range word1 {
		frequencies[c]++
	}

	for _, c := range word2 {
		frequencies[c]--
	}

	for _, num := range frequencies {
		if abs(num) > 3 {
			return false
		}
	}

	return true
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
