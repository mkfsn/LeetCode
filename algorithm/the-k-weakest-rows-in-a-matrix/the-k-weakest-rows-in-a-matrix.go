package the_k_weakest_rows_in_a_matrix

import (
	"sort"
)

type node struct {
	idx int
	row []int
}

func kWeakestRows(mat [][]int, k int) []int {
	nodes := make([]node, len(mat))

	for i, row := range mat {
		nodes[i] = node{idx: i, row: row}
	}

	countFn := func(nums []int) int {
		res := 0
		for _, n := range nums {
			if n == 0 {
				return res
			}
			res++
		}
		return res
	}

	sort.Slice(nodes, func(i, j int) bool {
		c1, c2 := countFn(nodes[i].row), countFn(nodes[j].row)

		if c1 == c2 {
			return nodes[i].idx < nodes[j].idx
		}

		return c1 < c2
	})

	res := make([]int, k)

	for i, n := range nodes[:k] {
		res[i] = n.idx
	}

	return res
}
