package query_kth_smallest_trimmed_number

import (
	"sort"
)

type item struct {
	index int
	value string
}

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {
	digits := len(nums[0])

	table := make([][]item, digits)
	base := 10

	for i := 0; i < digits; i++ {
		row := make([]item, 0)
		for j, v := range nums {
			row = append(row, item{index: j, value: v[digits-i-1:]})
		}

		sort.Slice(row, func(i, j int) bool {
			vi, vj := row[i].value, row[j].value
			if vi == vj {
				return row[i].index < row[j].index
			}
			return vi < vj
		})

		table[i] = row
		base *= 10
	}

	res := make([]int, len(queries))

	for i, query := range queries {
		res[i] = table[query[1]-1][query[0]-1].index
	}

	return res
}
