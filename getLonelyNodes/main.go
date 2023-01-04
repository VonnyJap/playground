package main

import "fmt"

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	fmt.Println(getLonelyNodes(root))

	root1 := &TreeNode{Val: 7}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 4}
	root1.Left.Left = &TreeNode{Val: 6}
	root1.Right.Left = &TreeNode{Val: 5}
	root1.Right.Right = &TreeNode{Val: 3}
	root1.Right.Right.Right = &TreeNode{Val: 2}

	fmt.Println(getLonelyNodes(root1))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLonelyNodes(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{}
	}

	result := []int{}

	if root.Left == nil {
		result = append(result, root.Right.Val)
	}
	if root.Right == nil {
		result = append(result, root.Left.Val)
	}

	result = append(result, getLonelyNodes(root.Left)...)
	result = append(result, getLonelyNodes(root.Right)...)

	return result
}
