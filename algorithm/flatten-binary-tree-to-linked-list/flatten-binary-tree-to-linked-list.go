package flatten_binary_tree_to_linked_list

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		tmp := findNode(root.Left)
		tmp.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	flatten(root.Right)
}

func findNode(root *TreeNode) *TreeNode {
	if root.Right == nil {
		return root
	}

	return findNode(root.Right)
}
