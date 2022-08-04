package clone_graph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNodes := make(map[int]*Node)
	traverse(node, newNodes)
	return newNodes[node.Val]
}

func traverse(node *Node, newNodes map[int]*Node) {
	if _, ok := newNodes[node.Val]; ok {
		return
	}

	newNode := &Node{Val: node.Val}
	newNodes[node.Val] = newNode

	for _, neighbor := range node.Neighbors {
		traverse(neighbor, newNodes)
	}

	for _, neighbor := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, newNodes[neighbor.Val])
	}
}

//  1 - 2
//  |   |
//  4 - 3

// 101
//   \
//    1 - 2
//    |   |
//    4 - 3

// 101     102
//   \     /
//    1 - 2
//    |   |
//    4 - 3

// 101     102
//   \     /
//    1 - 2
//    |   |
//    4 - 3
//   /
// 104

// 101     102
//   \     /
//    1 - 2
//    |   |
//    4 - 3
//   /     \
// 104      105
