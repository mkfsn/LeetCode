package find_and_replace_pattern

func findAndReplacePattern(words []string, pattern string) []string {
	encodedPattern := encode(pattern)
	res := make([]string, 0)
	for _, word := range words {
		if encodeEqually(encode(word), encodedPattern) {
			res = append(res, word)
		}
	}
	return res
}

func encode(s string) []int {
	dict := make(map[int32]int)
	counter := 1
	res := make([]int, 0, len(s))
	for _, c := range s {
		if v, ok := dict[c]; ok {
			res = append(res, v)
			continue
		}
		dict[c] = counter
		res = append(res, counter)
		counter++
	}
	return res
}

func encodeEqually(v1, v2 []int) bool {
	if len(v1) != len(v2) {
		return false
	}
	for i, v := range v1 {
		if v != v2[i] {
			return false
		}
	}
	return true
}
