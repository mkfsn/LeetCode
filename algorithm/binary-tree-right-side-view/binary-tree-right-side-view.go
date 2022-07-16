package binary_tree_right_side_view

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

func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)

	var traverse func(*TreeNode, int)

	traverse = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth >= len(res) {
			res = append(res, node.Val)
		} else {
			res[depth] = node.Val
		}

		if node.Left != nil {
			traverse(node.Left, depth+1)
		}

		if node.Right != nil {
			traverse(node.Right, depth+1)
		}
	}

	traverse(root, 0)

	return res
}
