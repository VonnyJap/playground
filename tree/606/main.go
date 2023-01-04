package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 4}
	// root.Left.Right = &TreeNode{Val: 5}
	// root.Left.Left.Right = &TreeNode{Val: 11}

	root.Right = &TreeNode{Val: 3}
	// root.Right.Left = &TreeNode{Val: 9}
	// root.Right.Right = &TreeNode{Val: 10}

	fmt.Println(tree2str(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// since it is preorder traverse
// step:
// 1. append the root to string
// 2. traverse right and left
// 3. if left empty - if right not empty then print "()", if right empty then dont do anything
// 4. if right not empty then do "()"

func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := fmt.Sprintf("%d", root.Val)
	left := tree2str(root.Left)
	right := tree2str(root.Right)
	if left != "" || (left == "" && right != "") {
		result = fmt.Sprintf("%s(%s)", result, left)
	}
	if right != "" {
		result = fmt.Sprintf("%s(%s)", result, right)
	}

	return result
}
