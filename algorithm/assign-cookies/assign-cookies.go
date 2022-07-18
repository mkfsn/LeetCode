package assign_cookies

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	cnt := 0

	for i, j := 0, 0; i < len(g) && j < len(s); {
		if g[i] <= s[j] {
			cnt++
			i++
			j++
		} else {
			j++
		}
	}

	return cnt
}

// 14:03
// 14:09
