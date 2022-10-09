package main

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	upsideDownBinaryTree(root)
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// The original left child becomes the new root.
// The original root becomes the new right child.
// The original right child becomes the new left child.
func upsideDownBinaryTree(root *TreeNode) *TreeNode {

	if root == nil || root.Left == nil {
		return root
	}

	var next *TreeNode = nil
	var prevRight *TreeNode = nil
	var prev *TreeNode = nil
	var current *TreeNode = root
	for {
		if current == nil {
			break
		}
		next = current.Left
		current.Left = prevRight
		prevRight = current.Right
		current.Right = prev

		prev = current
		current = next

	}

	root = prev
	return root

}
