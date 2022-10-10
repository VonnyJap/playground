package main

import "fmt"

func main() {

	root := &TreeNode{Val: 8}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 6}
	root.Left.Right.Left = &TreeNode{Val: 4}
	root.Left.Right.Right = &TreeNode{Val: 7}

	root.Right = &TreeNode{Val: 10}
	root.Right.Right = &TreeNode{Val: 14}
	root.Right.Right.Left = &TreeNode{Val: 13}

	fmt.Println(isValidBST(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	traverse := inorderTraversal(root)
	fmt.Println(traverse)
	for i := 0; i < len(traverse)-1; i++ {
		if traverse[i+1] <= traverse[i] {
			return false
		}
	}

	return true
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if (root.Left == nil) && (root.Right == nil) {
		return []int{root.Val}
	}
	var nodes []int
	nodes = append(nodes, inorderTraversal(root.Left)...)
	nodes = append(nodes, []int{root.Val}...)
	nodes = append(nodes, inorderTraversal(root.Right)...)
	return nodes
}

// what is bst:
// left -> less than the current
// right -> more than the current

// 1. traverse the tree using preorder
// 2. each traversal - default the check to true
// 3. if the children not nil then compare the value
// 	- return true when left smaller than root and subsequent recursive value
// 	- return false when right larger than root and also its subsequent recursive value
// 4. the recursion or traversal stop when root == nil or when root.Left == nil && root.Right == nil

// base case - when no children at all then return true
// if root.Left not nill - traverse it
// - when the last value of left is less than root.Val or when it is empty then set to true
// if root.Right not nill - traverse it
// - when the first value of right is bigger than root.Val or when it is empty then set to true
