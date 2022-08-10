package top_k_frequent_elements

import (
	"sort"
)

type node struct {
	num   int
	count int
}

func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	nodes := make([]node, 0, len(freq))

	for num, count := range freq {
		nodes = append(nodes, node{num, count})
	}

	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].count > nodes[j].count
	})

	res := make([]int, 0, k)

	for i := 0; i < k; i++ {
		res = append(res, nodes[i].num)
	}

	return res
}
