package n_ary_tree_preorder_traversal

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	result := make([]int, 0)
	dfs(root, &result)
	return result
}

func dfs(root *Node, result *[]int) {
	if root == nil {
		return
	}

	*result = append(*result, root.Val)

	for _, node := range root.Children {
		dfs(node, result)
	}
}
