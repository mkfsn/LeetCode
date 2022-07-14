package evaluate_boolean_binary_tree

import (
	"testing"
)

func Test_evaluateTree(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want bool
	}{
		{
			root: newTree([]interface{}{2, 1, 3, nil, nil, 0, 1}),
			want: true,
		},
		{
			root: newTree([]interface{}{1}),
			want: true,
		},
		{
			root: newTree([]interface{}{0}),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateTree(tt.root); got != tt.want {
				t.Errorf("evaluateTree() = %v, want %v", got, tt.want)
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
