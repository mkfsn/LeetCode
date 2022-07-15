package add_two_numbers

import (
	"fmt"
)

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	sum := 0
	cur := res

	for l1 != nil || l2 != nil || sum != 0 {
		sum /= 10

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		cur.Val += sum % 10

		// printList(res)

		if sum >= 10 || l1 != nil || l2 != nil {
			cur.Next = &ListNode{}
			cur = cur.Next
		}
	}

	return res
}

func printList(list *ListNode) {
	if list == nil {
		fmt.Println("x")
		return
	}

	fmt.Printf("%d -> ", list.Val)
	printList(list.Next)
}
