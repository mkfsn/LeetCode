package find_resultant_array_after_removing_anagrams

import (
	"sort"
)

func removeAnagrams(words []string) []string {

	seen := make(map[string]int)

	res := make([]string, 0)

	for i, word := range words {
		tmp := []byte(word)

		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})

		idx, ok := seen[string(tmp)]
		if ok && i == idx+1 {
			seen[string(tmp)] = i
			continue
		}
		seen[string(tmp)] = i

		res = append(res, word)
	}

	return res
}

// 11:08
// 11:24
