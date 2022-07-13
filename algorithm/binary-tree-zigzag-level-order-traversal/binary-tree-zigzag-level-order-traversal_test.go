package binary_tree_zigzag_level_order_traversal

import (
	"reflect"
	"testing"
)

func Test_zigzagLevelOrder(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want [][]int
	}{
		{
			root: newTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
			want: [][]int{
				{3},
				{20, 9},
				{15, 7},
			},
		},
		{
			root: newTree([]interface{}{1}),
			want: [][]int{
				{1},
			},
		},
		{
			root: newTree([]interface{}{}),
			want: [][]int{},
		},
		{
			root: newTree([]interface{}{1, 2, 3, 4, nil, nil, 5}),
			want: [][]int{
				{1},
				{3, 2},
				{4, 5},
			},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := zigzagLevelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zigzagLevelOrder() = %v, want %v", got, tt.want)
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

	if len(nodes) > 0 {
		return nodes[0]
	}

	return nil
}
