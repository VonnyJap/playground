package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// since this is a tree - traverse it
// example: 2 - 1 - 3 -> 2 - 3 - 1 -> this means left <-> right and root is the same
// base case -> nil will return nil
// base case -> no children - will return that nodes itself

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return root
	}

	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
