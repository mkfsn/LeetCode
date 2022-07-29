package valid_anagram

import (
	"sort"
)

func isAnagram(s string, t string) bool {
	x, y := []byte(s), []byte(t)

	sort.Slice(x, func(i, j int) bool {
		return x[i] < x[j]
	})

	sort.Slice(y, func(i, j int) bool {
		return y[i] < y[j]
	})

	return string(x) == string(y)
}
