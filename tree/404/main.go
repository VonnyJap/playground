package main

import "fmt"

func main() {

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	// fmt.Println(inorder(root))
	fmt.Println(sumOfLeftLeaves(root))
	fmt.Println(sumOfLeftLeavesRecur(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeavesRecur(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return processLeaves(root, false)
}

func processLeaves(root *TreeNode, isLeft bool) int {

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		if isLeft {
			return root.Val
		}
		return 0
	}
	total := 0
	if root.Left != nil {
		total += processLeaves(root.Left, true)
	}
	if root.Right != nil {
		total += processLeaves(root.Right, false)
	}
	return total
}

func sumOfLeftLeaves(root *TreeNode) int {

	if root == nil {
		return 0
	}

	stack := []*TreeNode{root}
	result := 0
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if isLeaf(current.Left) {
			result += current.Left.Val
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
	}

	return result
}

func isLeaf(root *TreeNode) bool {

	if root != nil && root.Left == nil && root.Right == nil {
		return true
	}
	return false
}

func inorder(root *TreeNode) []int {

	current := root
	stack := []*TreeNode{}
	result := []int{}
	for {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
			continue
		}
		if len(stack) > 0 {
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result = append(result, current.Val)
			current = current.Right
			continue
		}
		break
	}

	return result
}

// this is tree traverse again
// we can traverse all left and never include the right and root.Val
// and at the same time we only want the leaf

// 0 - 1 - 2
// when root == nil -> return 0
// when root.Left == nil && root.Right == nil -> return root.Val
// when sumOfLeftLeaves(root.Left)

// same thing use stack for this implementation
