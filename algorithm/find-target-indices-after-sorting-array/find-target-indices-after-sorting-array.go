package find_target_indices_after_sorting_array

import (
	"sort"
)

func targetIndices(nums []int, target int) []int {

	res := make([]int, 0)

	sort.Ints(nums)

	for i, v := range nums {
		if v == target {
			res = append(res, i)
		}
	}

	return res

}

// 11:05
// 11:08
