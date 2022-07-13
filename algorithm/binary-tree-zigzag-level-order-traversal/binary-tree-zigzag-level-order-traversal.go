package binary_tree_zigzag_level_order_traversal

import (
	"container/list"
)

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

func zigzagLevelOrder(root *TreeNode) [][]int {
	q := list.New()

	if root != nil {
		q.PushBack(root)
	}

	res := make([][]int, 0)

	leftToRight := true

	for q.Len() > 0 {
		values := make([]int, 0)

		n := q.Len()

		for i := 0; i < n; i++ {
			element := q.Front()
			q.Remove(element)

			node := element.Value.(*TreeNode)
			if node.Left != nil {
				q.PushBack(node.Left)
			}

			if node.Right != nil {
				q.PushBack(node.Right)
			}

			if leftToRight {
				values = append(values, node.Val)
			} else {
				values = append([]int{node.Val}, values...)
			}
		}

		leftToRight = !leftToRight

		res = append(res, values)
	}

	return res
}
