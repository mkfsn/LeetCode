package kth_smallest_element_in_a_bst

import (
	"testing"
)

func Test_kthSmallest(t *testing.T) {
	tests := []struct {
		root *TreeNode
		k    int
		want int
	}{
		{
			root: newTree([]interface{}{3, 1, 4, nil, 2}),
			k:    1,
			want: 1,
		},
		{
			root: newTree([]interface{}{5, 3, 6, 2, 4, nil, nil, 1}),
			k:    3,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := kthSmallest(tt.root, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
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

	for i, node := range nodes {
		left := 2*(i+1) - 1
		if left < len(nodes) && nodes[left] != nil {
			node.Left = nodes[left]
		}

		right := 2 * (i + 1)
		if right < len(nodes) && nodes[right] != nil {
			node.Right = nodes[right]
		}
	}

	return nodes[0]
}
