package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	// root.Right = &TreeNode{Val: 9}
	fmt.Println("hasPathSum: ", hasPathSum(&root, 1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// all add together becomes the targetSum
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
