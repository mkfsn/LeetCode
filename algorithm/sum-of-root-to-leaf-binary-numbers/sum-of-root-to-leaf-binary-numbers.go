package sum_of_root_to_leaf_binary_numbers

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

func sumRootToLeaf(root *TreeNode) int {
	return traverse(root, 0)
}

func traverse(root *TreeNode, base int) int {
	if root.Left == nil && root.Right == nil {
		return base<<1 + root.Val
	}

	var val int

	if root.Left != nil {
		val += traverse(root.Left, base<<1+root.Val)
	}

	if root.Right != nil {
		val += traverse(root.Right, base<<1+root.Val)
	}

	return val
}

// 11:56
// 12:00
