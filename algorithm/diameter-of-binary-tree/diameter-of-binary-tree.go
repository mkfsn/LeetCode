package diameter_of_binary_tree

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

func diameterOfBinaryTree(root *TreeNode) int {
	maxLength := 0
	traverse(root, &maxLength)
	return maxLength
}

func traverse(root *TreeNode, maxLength *int) int {
	if root == nil {
		return 0
	}

	len1 := traverse(root.Left, maxLength)
	len2 := traverse(root.Right, maxLength)
	*maxLength = getMax(*maxLength, len1+len2)
	return getMax(len1, len2) + 1
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
