package univalued_binary_tree

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

func isUnivalTree(root *TreeNode) bool {
	return traverse(root.Left, root.Val) && traverse(root.Right, root.Val)
}

func traverse(root *TreeNode, target int) bool {
	if root == nil {
		return true
	}

	if root.Val != target {
		return false
	}

	return traverse(root.Left, target) && traverse(root.Right, target)
}

// 12:01
// 12:03
