package middle_of_the_linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	count := 0

	cur := head
	for cur != nil {
		cur = cur.Next
		count++
	}

	count = count / 2

	cur = head
	for count > 0 {
		cur = cur.Next
		count--
	}

	return cur
}
