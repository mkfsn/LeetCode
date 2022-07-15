package verifying_an_alien_dictionary

func isAlienSorted(words []string, order string) bool {
	mapper := make(map[uint8]int)

	for i, ch := range order {
		mapper[uint8(ch)] = i
	}

	return isSorted(words, func(i, j int) bool {
		len1, len2 := len(words[i]), len(words[j])

		size := len1
		if len2 < size {
			size = len2
		}

		for k := 0; k < size; k++ {
			x, y := mapper[words[i][k]], mapper[words[j][k]]
			if x == y {
				continue
			} else if x < y {
				return true
			} else {
				return false
			}
		}

		return len1 < len2
	})
}

func isSorted(slice []string, less func(i, j int) bool) bool {
	n := len(slice)

	for i := n - 1; i > 0; i-- {
		if less(i, i-1) {
			return false
		}
	}

	return true
}
