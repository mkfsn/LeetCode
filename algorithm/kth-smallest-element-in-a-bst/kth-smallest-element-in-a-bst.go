package kth_smallest_element_in_a_bst

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

func kthSmallest(root *TreeNode, k int) int {
	v, _ := traverse(root, &k)
	return v
}

func traverse(root *TreeNode, k *int) (int, bool) {
	// fmt.Printf("[%d]root: %+v\n", *k, root)
	if root == nil {
		return -1, false
	}

	if v, ok := traverse(root.Left, k); ok {
		return v, ok
	}

	*k--
	if *k == 0 {
		return root.Val, true
	}

	return traverse(root.Right, k)
}
