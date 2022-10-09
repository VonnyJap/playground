package main

import "fmt"

func main() {
	root := TreeNode{Val: 0}
	root.Left = &TreeNode{Val: -10}
	root.Right = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: -3}
	root.Right.Right = &TreeNode{Val: 9}
	fmt.Println("min: ", minDepth(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := minDepth(root.Left)
	rh := minDepth(root.Right)
	if lh == 0 {
		return 1 + rh
	}
	if rh == 0 {
		return 1 + lh
	}
	if lh < rh {
		return 1 + lh
	}
	return 1 + rh
}
