package kth_smallest_element_in_a_sorted_matrix

import (
	"container/heap"
)

type item struct {
	value int

	row int
	col int
}

type MinHeap []*item

func (m MinHeap) Len() int {
	return len(m)
}

func (m MinHeap) Less(i, j int) bool {
	return m[i].value < m[j].value
}

func (m MinHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *MinHeap) Push(x interface{}) {
	*m = append(*m, x.(*item))
}

func (m *MinHeap) Pop() interface{} {
	n := len(*m)
	item := (*m)[n-1]
	*m = (*m)[0 : n-1]
	return item
}

func NewMinHeap(n int) heap.Interface {
	res := make(MinHeap, n)
	return &res
}

func kthSmallest(matrix [][]int, k int) int {
	rows, cols := len(matrix), len(matrix[0])

	pq := NewMinHeap(0)

	heap.Init(pq)

	for i := 0; i < rows; i++ {
		heap.Push(pq, &item{
			value: matrix[i][0],
			row:   i,
			col:   0,
		})
	}

	res := 0
	for i := 0; i < k; i++ {
		x := heap.Pop(pq).(*item)
		res = x.value
		if x.col+1 < cols {
			heap.Push(pq, &item{
				value: matrix[x.row][x.col+1],
				row:   x.row,
				col:   x.col + 1,
			})
		}
	}

	return res
}
