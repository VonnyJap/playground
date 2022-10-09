package main

import "fmt"

func main() {
	five := TreeNode{Val: 5}
	six := TreeNode{Val: 6}
	nine := TreeNode{Val: 9}
	twelve := TreeNode{
		Val:   12,
		Left:  &five,
		Right: &six,
	}
	one := TreeNode{
		Val:   1,
		Left:  &twelve,
		Right: &nine,
	}
	fmt.Println(inorderTraversal(&one))
}

// inorder traveral
// left then root then right

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if (root.Left == nil) && (root.Right == nil) {
		return []int{root.Val}
	}
	var nodes []int
	nodes = append(nodes, inorderTraversal(root.Left)...)
	nodes = append(nodes, []int{root.Val}...)
	nodes = append(nodes, inorderTraversal(root.Right)...)
	return nodes
}

// For traversing a (non-empty) binary tree in an inorder fashion, we must do these three things for every node n starting from the tree’s root:

// (L) Recursively traverse its left subtree. When this step is finished, we are back at n again.
// (N) Process n itself.
// (R) Recursively traverse its right subtree. When this step is finished, we are back at n again.
