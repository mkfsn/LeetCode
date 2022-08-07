package linked_list_cycle_ii

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

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head

	for fast != nil && slow != nil {
		slow = slow.Next

		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next

		if slow == fast {
			// find the beginning of the cycle

			for head != slow {
				slow = slow.Next
				head = head.Next
			}

			return head

		}
	}

	return nil
}
