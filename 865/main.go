package main

import "fmt"

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	node := subtreeWithAllDeepest(root)

	fmt.Printf("%#v\n", node)
	fmt.Printf("%#v\n", node.Left)
	fmt.Printf("%#v\n", node.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	_, node := find(root, 0)
	return node
}

func find(node *TreeNode, level int) (int, *TreeNode) {

	if node == nil {
		return 0, nil
	}
	if node.Left == nil && node.Right == nil {
		return level + 1, node
	}
	left, leftNode := find(node.Left, level+1)
	right, rightNode := find(node.Right, level+1)

	if left == right {
		return left, node
	}
	if left > right {
		return left, leftNode
	}
	return right, rightNode
}

// Approach: Follow the steps below to solve the problem:

// Traverse the  Binary Tree recursively using  DFS .

// For every node, find the depth of its left and right subtrees.

// If depth of the left subtree > depth of the right subtree: Traverse the left subtree.

// If depth of the right subtree > depth of the left subtree: Traverse the right subtree.

// Otherwise, return the current node.
