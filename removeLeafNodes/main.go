package main

import "fmt"

func main() {
	// root := &TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	// root.Left.Left = &TreeNode{Val: 2}
	// root.Right = &TreeNode{Val: 3}
	// root.Right.Left = &TreeNode{Val: 2}
	// root.Right.Right = &TreeNode{Val: 4}

	// chopped := removeLeafNodes(root, 2)
	// fmt.Printf("%+v\n", chopped)
	// fmt.Printf("%+v\n", chopped.Left)
	// fmt.Printf("%+v\n", chopped.Right)
	// fmt.Printf("%+v\n", chopped.Right.Left)
	// fmt.Printf("%+v\n", chopped.Right.Right)

	root1 := &TreeNode{Val: 1}
	root1.Left = &TreeNode{Val: 3}
	root1.Left.Left = &TreeNode{Val: 3}
	root1.Left.Right = &TreeNode{Val: 2}
	root1.Right = &TreeNode{Val: 3}

	chopped1 := removeLeafNodes(root1, 3)
	fmt.Printf("%+v\n", chopped1)
	fmt.Printf("%+v\n", chopped1.Left)
	fmt.Printf("%+v\n", chopped1.Right)
	fmt.Printf("%+v\n", chopped1.Left.Left)
	fmt.Printf("%+v\n", chopped1.Left.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {

	if root == nil {
		return root
	}

	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Left == nil && root.Right == nil {
		if root.Val == target {
			return nil
		}
		return root
	}

	return root
}
