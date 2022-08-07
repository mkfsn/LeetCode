package top_k_frequent_words

import (
	"container/heap"
)

type pair struct {
	word  string
	count int
}

type maxHeap []pair

func (m maxHeap) Len() int {
	return len(m)
}

// max
func (m maxHeap) Less(i, j int) bool {
	if m[i].count == m[j].count {
		return m[i].word < m[j].word
	}
	return m[i].count > m[j].count
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(pair))
}

func (m *maxHeap) Pop() interface{} {
	tmp := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return tmp
}

func newMaxHeap() heap.Interface {
	res := make(maxHeap, 0)
	return &res
}

func topKFrequent(words []string, k int) []string {
	dict := make(map[string]int)

	for _, word := range words {
		dict[word]++
	}

	h := newMaxHeap()
	heap.Init(h)

	for word, count := range dict {
		heap.Push(h, pair{word: word, count: count})
	}

	results := make([]string, 0, k)

	// pop k times
	for i := 0; i < k; i++ {
		tmp := heap.Pop(h).(pair)
		results = append(results, tmp.word)
	}

	return results
}
