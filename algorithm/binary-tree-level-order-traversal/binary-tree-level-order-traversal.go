package binary_tree_level_order_traversal

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

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	dfs(root, 1, &result)
	return result
}

func dfs(root *TreeNode, level int, result *[][]int) {
	if root == nil {
		return
	}

	if len(*result) < level {
		*result = append(*result, make([]int, 0))
	}

	arr := (*result)[level-1]
	(*result)[level-1] = append(arr, root.Val)

	dfs(root.Left, level+1, result)
	dfs(root.Right, level+1, result)
}
