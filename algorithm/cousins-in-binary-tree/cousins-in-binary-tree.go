package cousins_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isCousins(root *TreeNode, x int, y int) bool {
	depthMap := make(map[int]int)
	parentMap := make(map[int]int)

	traverse(root, 0, -1, depthMap, parentMap)

	return depthMap[x] == depthMap[y] && parentMap[x] != parentMap[y]
}

func traverse(node *TreeNode, depth int, parent int, depthMap map[int]int, parentMap map[int]int) {
	if node == nil {
		return
	}

	depthMap[node.Val] = depth
	parentMap[node.Val] = parent

	traverse(node.Right, depth+1, node.Val, depthMap, parentMap)
	traverse(node.Left, depth+1, node.Val, depthMap, parentMap)
}
