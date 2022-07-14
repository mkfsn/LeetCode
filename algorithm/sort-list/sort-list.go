package sort_list

import (
	"sort"
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

func sortList(head *ListNode) *ListNode {
	list := make([]*ListNode, 0)

	next := head
	for next != nil {
		list = append(list, next)
		next = next.Next
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].Val < list[j].Val
	})

	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}
	list[len(list)-1].Next = nil

	return list[0]
}
