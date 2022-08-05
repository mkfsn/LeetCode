package isomorphic_strings

func isIsomorphic(s string, t string) bool {
	return equalSlice(encode(s), encode(t))
}

func equalSlice(n1, n2 []int) bool {
	size := len(n1)

	for i := 0; i < size; i++ {
		if n1[i] != n2[i] {
			return false
		}
	}

	return true
}

func encode(s string) []int {
	idx := 0
	dict := make(map[int32]int)
	res := make([]int, 0, len(s))

	for _, c := range s {
		v, ok := dict[c]
		if ok {
			res = append(res, v)
			continue
		}

		dict[c] = idx
		res = append(res, idx)
		idx++
	}

	return res
}
