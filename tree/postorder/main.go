package main

import "fmt"

func main() {
	// root := &TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	// root.Left.Left = &TreeNode{Val: 4}
	// root.Left.Right = &TreeNode{Val: 5}
	// root.Right = &TreeNode{Val: 3}

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Left.Right = &TreeNode{Val: 11}

	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 10}

	fmt.Println(postorderTraversal(root))
}

// with the pre - we start from root

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var nodes []int
	if root.Left != nil {
		nodes = append(nodes, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nodes = append(nodes, postorderTraversal(root.Right)...)
	}
	nodes = append(nodes, []int{root.Val}...)
	return nodes
}
