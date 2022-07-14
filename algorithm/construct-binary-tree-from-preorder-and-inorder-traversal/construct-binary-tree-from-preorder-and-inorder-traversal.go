package construct_binary_tree_from_preorder_and_inorder_traversal

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	val := preorder[0]

	node := &TreeNode{Val: val}

	left, right := splitInOrder(val, inorder)

	if len(left) > 0 {
		node.Left = buildTree(preorder[1:1+len(left)], left)
	}

	if len(right) > 0 {
		node.Right = buildTree(preorder[1+len(left):], right)
	}

	return node
}

func splitInOrder(val int, inorder []int) ([]int, []int) {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == val {
			return inorder[:i], inorder[i+1:]
		}
	}

	return nil, nil
}
