package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	lh := maxDepth(root.Left)
	rh := maxDepth(root.Right)

	if math.Abs(float64(lh)-float64(rh)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

func main() {
	root := TreeNode{Val: 0}
	root.Left = &TreeNode{Val: -10}
	root.Right = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: -3}
	root.Right.Right = &TreeNode{Val: 9}

	if isBalanced(&root) {
		fmt.Println("Tree is balanced")
	} else {
		fmt.Println("Tree is not balanced")
	}
}
