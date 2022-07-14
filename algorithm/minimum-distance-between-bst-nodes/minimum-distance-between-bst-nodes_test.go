package minimum_distance_between_bst_nodes

import (
	"container/list"
	"testing"
)

func Test_minDiffInBST(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			root: newTree([]interface{}{4, 2, 6, 1, 3}),
			want: 1,
		},
		{
			root: newTree([]interface{}{1, 0, 48, nil, nil, 12, 49}),
			want: 1,
		},
		{
			root: newTree([]interface{}{3, 1}),
			want: 2,
		},
		{
			root: newTree([]interface{}{27, nil, 34, nil, 58, 50, nil, 44}),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDiffInBST(tt.root); got != tt.want {
				t.Errorf("minDiffInBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTree(arr []interface{}) *TreeNode {
	nodes := make([]*TreeNode, len(arr))

	for i, v := range arr {
		switch v.(type) {
		case nil:
			nodes[i] = nil
		case int:
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}

	if len(nodes) < 1 {
		return nil
	}

	q := list.New()

	q.PushBack(nodes[0])
	idx := 1

	for q.Len() > 0 {
		size := q.Len()
		for i := 0; i < size; i++ {
			e := q.Front()
			q.Remove(e)

			node := e.Value.(*TreeNode)

			if idx < len(nodes) {
				node.Left = nodes[idx]
			}

			if idx+1 < len(nodes) {
				node.Right = nodes[idx+1]
			}

			idx += 2

			if node.Left != nil {
				q.PushBack(node.Left)
			}

			if node.Right != nil {
				q.PushBack(node.Right)
			}
		}
	}

	return nodes[0]
}
