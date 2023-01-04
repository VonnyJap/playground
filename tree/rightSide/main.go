package main

import "fmt"

func main() {
	// root := &TreeNode{}
	// root.Left = &TreeNode{Val: 2}
	// root.Left.Left = &TreeNode{Val: 4}
	// root.Left.Right = &TreeNode{Val: 5}
	// root.Left.Left.Right = &TreeNode{Val: 11}

	// root.Right = &TreeNode{Val: 3}
	// root.Right.Left = &TreeNode{Val: 9}
	// root.Right.Right = &TreeNode{Val: 4}

	fmt.Println(rightSideView(nil))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	nodes := preorderWithLevel(root)
	for _, n := range nodes {
		result = append(result, n[len(n)-1])
	}
	return result
}

func preorderWithLevel(root *TreeNode) [][]int {

	type TreeNodeWithLevel struct {
		*TreeNode
		Level int
	}
	stack := []*TreeNodeWithLevel{{root, 0}}
	current := &TreeNodeWithLevel{}
	result := [][]int{}

	for len(stack) > 0 {
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(result) > current.Level {
			result[current.Level] = append(result[current.Level], current.Val)
		} else {
			result = append(result, []int{current.Val})
		}
		if current.Right != nil {
			stack = append(stack, &TreeNodeWithLevel{current.Right, current.Level + 1})
		}
		if current.Left != nil {
			stack = append(stack, &TreeNodeWithLevel{current.Left, current.Level + 1})
		}
	}

	return result
}
