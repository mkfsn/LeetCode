package last_stone_weight

import (
	"container/heap"
)

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() interface{} {
	n := len(*m)
	v := (*m)[n-1]
	*m = (*m)[:n-1]
	return v
}

func newMinHeap() heap.Interface {
	res := make(maxHeap, 0)
	return &res
}

func lastStoneWeight(stones []int) int {
	h := newMinHeap()
	heap.Init(h)

	for _, n := range stones {
		heap.Push(h, n)
	}

	for h.Len() >= 2 {
		a := heap.Pop(h).(int)
		b := heap.Pop(h).(int)

		if a != b {
			heap.Push(h, a-b)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return heap.Pop(h).(int)
}
