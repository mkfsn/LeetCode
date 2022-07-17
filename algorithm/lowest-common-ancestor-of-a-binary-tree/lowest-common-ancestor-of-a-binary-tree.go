package lowest_common_ancestor_of_a_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pathP := traverse(root, p)
	pathQ := traverse(root, q)

	paths := make(map[*TreeNode]bool)

	for _, path := range pathP {
		paths[path] = true
	}

	for _, path := range pathQ {
		if paths[path] {
			return path
		}
	}

	return nil
}

func traverse(root *TreeNode, target *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	if root == target {
		return []*TreeNode{root}
	}

	if left := traverse(root.Left, target); left != nil {
		return append(left, root)
		// return append([]*TreeNode{root}, left...)
	}

	if right := traverse(root.Right, target); right != nil {
		return append(right, root)
		// return append([]*TreeNode{root}, right...)
	}

	// Not found
	return nil
}
