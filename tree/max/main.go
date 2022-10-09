package main

import "fmt"

func main() {
	// root = [3,9,20,null,null,15,7]
	// nine := TreeNode{Val: 9}
	// fifteen := TreeNode{Val: 15}
	// seven := TreeNode{Val: 7}
	twenty := TreeNode{Val: 20}
	three := TreeNode{
		Val:   3,
		Right: &twenty,
	}
	fmt.Println(maxDepth(&three))
}

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
