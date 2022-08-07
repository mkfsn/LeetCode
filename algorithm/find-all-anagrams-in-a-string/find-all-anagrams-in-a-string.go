package find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	countP := make(map[uint8]int)
	for i := range p {
		countP[p[i]]++
	}

	countS := make(map[uint8]int)
	for i := 0; i < len(s) && i < len(p)-1; i++ {
		countS[s[i]]++
	}

	result := make([]int, 0)

	for i := len(p) - 1; i < len(s); i++ {
		countS[s[i]]++

		if equalMap(countP, countS) {
			result = append(result, i-len(p)+1)
		}

		countS[s[i-len(p)+1]]--
	}

	return result
}

func equalMap(m1, m2 map[uint8]int) bool {
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
