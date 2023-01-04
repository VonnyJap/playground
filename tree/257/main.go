package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	fmt.Println(binaryTreePaths(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return []string{}
	}
	// when we reach here - we know that the root is no longer nil and should have some val
	result := []string{}
	leftSide := binaryTreePaths(root.Left)
	rightSide := binaryTreePaths(root.Right)
	if len(leftSide) == 0 && len(rightSide) == 0 {
		return []string{fmt.Sprintf("%d", root.Val)}
	}
	for _, ls := range leftSide {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, ls))
	}
	for _, rs := range rightSide {
		result = append(result, fmt.Sprintf("%d->%s", root.Val, rs))
	}
	return result
}

// since this is a tree
// we can do traveral
// 1. pre
// 2. post
// 3. inorder

// steps:
// 1. traverse left all the way
// 2. traverse right all the way
// 3. combine root to the left and right
// 4. return all things together
// 5. stop condition is when we find nothing?
