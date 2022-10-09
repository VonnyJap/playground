package main

import "fmt"

func main() {
	// [2,1,4,3,null,5]
	root := &TreeNode{Val: 1}
	// root.Left = &TreeNode{Val: 2}
	// root.Right = &TreeNode{Val: 3}
	// root.Left.Left = &TreeNode{Val: 4}
	// root.Left.Right = &TreeNode{Val: 5}
	fmt.Println(findLeaves(root))
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var nodes []int
	if root.Left != nil {
		nodes = append(nodes, postorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nodes = append(nodes, postorderTraversal(root.Right)...)
	}
	nodes = append(nodes, []int{root.Val}...)
	return nodes
}

// if we use postorder then we can do left, right, root
// postorder means root the last

// what will happen here when we recurse thing
func findLeaves(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}

	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}

	left := findLeaves(root.Left)
	fmt.Println("left:", left)
	right := findLeaves(root.Right)
	fmt.Println("right: ", right)
	nodes := [][]int{}

	i := 0
	j := 0

	// each time we actually need to separate combine
	for i < len(left) && j < len(right) {
		combine := []int{}
		combine = append(combine, left[i]...)
		i++
		combine = append(combine, right[j]...)
		j++
		nodes = append(nodes, [][]int{combine}...)
	}

	// exhaust the left if any
	for i < len(left) {
		nodes = append(nodes, [][]int{left[i]}...)
		i++
	}

	// exhaust the right if any
	for j < len(right) {
		nodes = append(nodes, [][]int{right[j]}...)
		j++
	}
	return append(nodes, [][]int{{root.Val}}...)
}
