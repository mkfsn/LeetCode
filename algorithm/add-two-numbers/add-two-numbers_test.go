package add_two_numbers

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			l1:   newList(2, 4, 3),
			l2:   newList(5, 6, 4),
			want: newList(7, 0, 8),
		},
		{
			l1:   newList(0),
			l2:   newList(0),
			want: newList(0),
		},
		{
			l1:   newList(9, 9, 9, 9, 9, 9, 9),
			l2:   newList(9, 9, 9, 9),
			want: newList(8, 9, 9, 9, 0, 0, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !equalList(got, tt.want) {
				t.Errorf("got:")
				printList(got)
				t.Errorf("want:")
				printList(tt.want)
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
