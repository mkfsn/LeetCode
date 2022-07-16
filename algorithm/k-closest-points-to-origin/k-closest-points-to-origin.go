package k_closest_points_to_origin

import (
	"container/heap"
)

// return k smallest distance

type point struct {
	distance uint // skip the square part
	x, y     int
}

type minHeap []*point

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(i, j int) bool {
	return m[i].distance > m[j].distance
}

func (m minHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(*point))
}

func (m *minHeap) Pop() interface{} {
	item := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return item
}

func newMinHeap() heap.Interface {
	res := make(minHeap, 0)
	return &res
}

func kClosest(points [][]int, k int) [][]int {
	pq := newMinHeap()
	heap.Init(pq)

	for _, p := range points {
		if pq.Len() < k {
			heap.Push(pq, &point{
				distance: uint(p[0]*p[0]) + uint(p[1]*p[1]),
				x:        p[0],
				y:        p[1],
			})
			continue
		}

		biggest := heap.Pop(pq).(*point)
		newPoint := &point{
			distance: uint(p[0]*p[0]) + uint(p[1]*p[1]),
			x:        p[0],
			y:        p[1],
		}

		if newPoint.distance < biggest.distance {
			heap.Push(pq, newPoint)
		} else {
			heap.Push(pq, biggest)
		}
	}

	res := make([][]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(pq).(*point)
		res[i] = []int{p.x, p.y}
	}
	return res
}
