package validate_binary_search_tree

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

func isValidBST(root *TreeNode) bool {
	_, _, ok := isValidBSTTraversal(root)
	return ok
}

func isValidBSTTraversal(root *TreeNode) (int, int, bool) {
	var lMin, lMax, rMin, rMax int
	var lValid, rValid bool

	if root.Right == nil && root.Left == nil {
		return root.Val, root.Val, true
	}

	if root.Left != nil {
		lMin, lMax, lValid = isValidBSTTraversal(root.Left)
		if !lValid {
			return -1, -1, false
		}

		if lMax >= root.Val {
			return -1, -1, false
		}
	}

	if root.Right != nil {
		rMin, rMax, rValid = isValidBSTTraversal(root.Right)
		if !rValid {
			return -1, -1, false
		}
		if rMin <= root.Val {
			return -1, -1, false
		}
	}

	if lValid && !rValid {
		return lMin, getMax(lMax, root.Val), true
	}

	if rValid && !lValid {
		return getMin(rMin, root.Val), rMax, true
	}

	min, max := root.Val, root.Val

	min = getMin(min, lMin)
	max = getMax(max, rMax)

	return min, max, true
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
