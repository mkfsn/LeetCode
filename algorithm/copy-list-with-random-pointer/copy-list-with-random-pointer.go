package copy_list_with_random_pointer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}

	cur := head
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur.Next = newNode
		cur = newNode.Next
	}

	cur = head
	for cur != nil {
		newNode := cur.Next
		if newNode != nil && cur.Random != nil {
			newNode.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	newHead := head.Next

	cur = head
	newNode := head.Next
	for cur != nil {
		cur.Next = cur.Next.Next
		cur = cur.Next

		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
			newNode = newNode.Next
		}
	}

	return newHead
}
