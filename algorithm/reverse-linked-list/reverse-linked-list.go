package reverse_linked_list

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

func reverseList(head *ListNode) *ListNode {
	tmp := ListNode{Next: nil}

	for head != nil {
		t := head
		head = head.Next

		t.Next = tmp.Next
		tmp.Next = t
	}

	return tmp.Next
}
