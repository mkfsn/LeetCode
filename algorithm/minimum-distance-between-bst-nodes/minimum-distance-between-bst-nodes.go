package minimum_distance_between_bst_nodes

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

func minDiffInBST(root *TreeNode) int {
	res := -1

	if root.Left != nil {
		res = root.Val - root.Left.Val
	} else if root.Right != nil {
		res = root.Right.Val - root.Val
	}

	prev := (*TreeNode)(nil)

	var traverse func(root *TreeNode)

	// inorder traverse
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		if root.Left != nil {
			traverse(root.Left)
		}

		if prev != nil {
			res = getMin(res, root.Val-prev.Val)
		}

		// fmt.Printf("root: %v, prev: %v\n", root, prev)
		prev = root

		if root.Right != nil {
			traverse(root.Right)
		}
	}

	traverse(root)
	return res

}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
