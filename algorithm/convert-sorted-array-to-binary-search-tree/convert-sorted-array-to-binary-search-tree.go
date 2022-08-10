package convert_sorted_array_to_binary_search_tree

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

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	// size >= 2, e.g.
	// - [3, 4]
	// - [3, 4, 5]

	mid := len(nums) / 2
	// - [3, 4] -> 1
	// - [3, 4, 5] -> 1

	left := nums[:mid]
	// [3]
	// [3]

	right := nums[mid+1:]
	// []
	// [5]

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(left),
		Right: sortedArrayToBST(right),
	}
}
