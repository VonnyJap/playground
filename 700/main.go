package main

import "fmt"

func main() {

	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 7}

	fmt.Println(searchBST(root, 3))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {

	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	var Left *TreeNode
	var Right *TreeNode
	if root.Left != nil {
		Left = searchBST(root.Left, val)
	}
	if root.Right != nil {
		Right = searchBST(root.Right, val)
	}
	if Left != nil {
		return Left
	}
	return Right
}
