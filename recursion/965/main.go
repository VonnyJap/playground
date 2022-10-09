package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 1}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println(isUnivalTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	var same bool = true
	if root.Left != nil && root.Val != root.Left.Val {
		return false
	}
	if root.Right != nil && root.Val != root.Right.Val {
		return false
	}
	return same && isUnivalTree(root.Left) && isUnivalTree(root.Right)
}
