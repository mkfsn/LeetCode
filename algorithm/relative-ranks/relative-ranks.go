package relative_ranks

import (
	"fmt"
	"sort"
)

func findRelativeRanks(score []int) []string {

	sorted := make([]int, len(score))
	for i, v := range score {
		sorted[i] = v
	}
	sort.Ints(sorted)

	mapper := make(map[int]string)
	idx := 1

	for i := len(sorted) - 1; i >= 0; i-- {
		switch idx {
		case 1:
			mapper[sorted[i]] = "Gold Medal"
		case 2:
			mapper[sorted[i]] = "Silver Medal"
		case 3:
			mapper[sorted[i]] = "Bronze Medal"
		default:
			mapper[sorted[i]] = fmt.Sprintf("%d", idx)
		}

		idx++
	}

	res := make([]string, len(score))

	for i, n := range score {
		res[i] = mapper[n]
	}

	return res
}

// 14:11
// 14:17
