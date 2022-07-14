package construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	tests := []struct {
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			preorder: []int{3, 9, 20, 15, 7},
			inorder:  []int{9, 3, 15, 20, 7},
			want:     newTree([]interface{}{3, 9, 20, nil, nil, 15, 7}),
		},
		{
			preorder: []int{-1},
			inorder:  []int{-1},
			want:     newTree([]interface{}{-1}),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := buildTree(tt.preorder, tt.inorder); !equalTree(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func equalTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 != nil {
		return false
	} else if t1 != nil && t2 == nil {
		return false
	}

	if t1 == nil && t2 == nil {
		return true
	}

	if t1.Val != t2.Val {
		return false
	}

	return equalTree(t1.Left, t2.Left) && equalTree(t1.Right, t2.Right)
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
