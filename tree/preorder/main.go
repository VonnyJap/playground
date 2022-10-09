package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 3}

	fmt.Println(preorderTraversal(root))
}

// with the pre - we start from root

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var nodes []int
	nodes = append(nodes, []int{root.Val}...)
	if root.Left != nil {
		nodes = append(nodes, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nodes = append(nodes, preorderTraversal(root.Right)...)
	}
	return nodes
}
