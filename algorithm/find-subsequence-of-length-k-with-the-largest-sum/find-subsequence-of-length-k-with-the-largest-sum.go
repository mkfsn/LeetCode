package find_subsequence_of_length_k_with_the_largest_sum

import (
	"sort"
)

type node struct {
	idx int
	val int
}

func maxSubsequence(nums []int, k int) []int {
	nodes := make([]node, len(nums))
	for i, v := range nums {
		nodes[i] = node{idx: i, val: v}
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].val < nodes[j].val
	})

	sortedNodes := nodes[len(nodes)-k:]

	sort.Slice(sortedNodes, func(i, j int) bool {
		return sortedNodes[i].idx < sortedNodes[j].idx
	})

	res := make([]int, k)

	for i, v := range sortedNodes {
		res[i] = v.val
	}

	return res
}
