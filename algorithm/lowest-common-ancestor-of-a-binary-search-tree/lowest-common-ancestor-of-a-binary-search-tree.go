package lowest_common_ancestor_of_a_binary_search_tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for (p.Val-root.Val)*(q.Val-root.Val) > 0 {
		if p.Val > root.Val {
			root = root.Right
		} else {
			root = root.Left
		}
	}

	return root
}
