package evaluate_boolean_binary_tree

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

func evaluateTree(root *TreeNode) bool {
	switch root.Val {
	case 0: // false
		return false
	case 1: // true
		return true
	case 2: // or
		return evaluateTree(root.Left) || evaluateTree(root.Right)
	case 3: // and
		return evaluateTree(root.Left) && evaluateTree(root.Right)
	}

	return false // won't happen
}
