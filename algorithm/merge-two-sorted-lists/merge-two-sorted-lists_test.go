package merge_two_sorted_lists

import (
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			list1: newList(1, 2, 4),
			list2: newList(1, 3, 4),
			want:  newList(1, 1, 2, 3, 4, 4),
		},
		{
			list1: newList(1, 1, 3, 3, 5, 5, 9),
			list2: newList(1, 1, 2, 2, 6, 7, 10),
			want:  newList(1, 1, 1, 1, 2, 2, 3, 3, 5, 5, 6, 7, 9, 10),
		},
		{
			list1: newList(),
			list2: newList(),
			want:  newList(),
		},
		{
			list1: newList(),
			list2: newList(0),
			want:  newList(0),
		},
		{
			list1: newList(1),
			list2: newList(),
			want:  newList(1),
		},
		{
			list1: newList(2),
			list2: newList(1),
			want:  newList(1, 2),
		},
		{
			list1: newList(-1),
			list2: newList(1),
			want:  newList(-1, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.list1, tt.list2); !equalList(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newList(values ...int) *ListNode {
	list := make([]*ListNode, len(values))

	for i, v := range values {
		list[i] = &ListNode{Val: v}
	}

	for i := 0; i < len(values)-1; i++ {
		list[i].Next = list[i+1]
	}

	if len(list) == 0 {
		return nil
	}

	return list[0]
}

func equalList(l1, l2 *ListNode) bool {
	if l1 == nil && l2 != nil {
		return false
	} else if l1 != nil && l2 == nil {
		return false
	}

	if l1 != nil && l2 != nil {
		return l1.Val == l2.Val && equalList(l1.Next, l2.Next)
	}

	return true
}
