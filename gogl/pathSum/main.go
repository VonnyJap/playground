package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Left.Left = &TreeNode{Val: 11}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Left.Left.Right = &TreeNode{Val: 2}

	root.Right = &TreeNode{Val: 8}
	root.Right.Left = &TreeNode{Val: 13}
	root.Right.Right = &TreeNode{Val: 4}
	root.Right.Right.Right = &TreeNode{Val: 1}

	fmt.Println(pathSum(root, 22))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) bool {

	result := sum(root)

	for _, r := range result {
		if r == target {
			return true
		}
	}
	return false
}

func sum(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	result := []int{}
	for _, l := range sum(root.Left) {
		result = append(result, l+root.Val)
	}
	for _, r := range sum(root.Right) {
		result = append(result, r+root.Val)
	}
	return result
}
