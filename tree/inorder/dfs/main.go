package main

import "fmt"

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Left.Left.Right = &TreeNode{Val: 11}

	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 10}

	target := findTarget(root, 2)
	fmt.Println(preorder(target))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// get particular target root and then do another dfs from it?
// do preorder traverse then

func findTarget(root *TreeNode, targetVal int) *TreeNode {

	stack := []*TreeNode{root}
	current := &TreeNode{}

	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if current.Val == targetVal {
			break
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return current
}

func preorder(root *TreeNode) []int {

	stack := []*TreeNode{root}
	current := &TreeNode{}
	result := []int{}

	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
	}

	return result
}
