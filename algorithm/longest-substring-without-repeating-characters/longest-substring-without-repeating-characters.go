package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	lastIndex := make(map[int32]int)
	start := -1
	res := 0

	for i, c := range s {
		index, ok := lastIndex[c]
		if ok && index > start {
			start = index
		}

		lastIndex[c] = i
		res = getMax(res, i-start)
	}

	return res
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
