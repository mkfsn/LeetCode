package all_nodes_distance_k_in_binary_tree

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

// The number of nodes in the tree is in the range [1, 500].
// 0 <= Node.val <= 500
// All the values Node.val are unique.
// target is the value of one of the nodes in the tree.
// 0 <= k <= 1000
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parentTable := make(map[*TreeNode]*TreeNode)
	traverse(root, parentTable, nil)

	seen := make(map[*TreeNode]bool)
	res := make([]int, 0)
	var dfs func(root *TreeNode, k int)

	dfs = func(root *TreeNode, k int) {
		if root == nil {
			return
		}

		if seen[root] {
			return
		}
		seen[root] = true

		if k == 0 {
			res = append(res, root.Val)
		}

		dfs(root.Left, k-1)
		dfs(root.Right, k-1)
		dfs(parentTable[root], k-1)
	}

	dfs(target, k)

	return res
}

func traverse(root *TreeNode, parentTable map[*TreeNode]*TreeNode, parent *TreeNode) {
	if root == nil {
		return
	}
	parentTable[root] = parent
	traverse(root.Left, parentTable, root)
	traverse(root.Right, parentTable, root)
}
