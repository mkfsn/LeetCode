package minimum_absolute_difference_in_bst

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
func getMinimumDifference(root *TreeNode) int {
	var list listNode
	traverse(root, &list)

	min := -1

	cur := &list
	for cur != nil {
		if cur.Prev == nil {
			break
		}

		diff := abs(cur.Val - cur.Prev.Val)

		if min == -1 {
			min = diff
		} else if min > diff {
			min = diff
		}

		cur = cur.Prev
	}

	cur = &list
	for cur != nil {
		if cur.Next == nil {
			break
		}

		diff := abs(cur.Val - cur.Next.Val)

		if min == -1 {
			min = diff
		} else if min > diff {
			min = diff
		}

		cur = cur.Next
	}

	return min
}

type listNode struct {
	Prev *listNode
	Val  int
	Next *listNode
}

func traverse(root *TreeNode, list *listNode) {
	list.Val = root.Val

	if root.Left != nil {
		newNode := listNode{Next: list}
		if list.Prev != nil {
			list.Prev.Next = &newNode
			newNode.Prev = list.Prev
		}
		list.Prev = &newNode
		traverse(root.Left, &newNode)
	}

	if root.Right != nil {
		newNode := listNode{Prev: list}
		if list.Next != nil {
			list.Next.Prev = &newNode
			newNode.Next = list.Next
		}
		list.Next = &newNode
		traverse(root.Right, &newNode)
	}
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
