package main

import (
	"fmt"
	"math"
)

func main() {

	// root = [4,2,6,1,3]

	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 6}

	fmt.Println(getMinimumDifference(root))

	root = &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 48}
	root.Right.Left = &TreeNode{Val: 12}
	root.Right.Right = &TreeNode{Val: 49}

	fmt.Println(getMinimumDifference(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {

	nodes := inorderTraversal(root)

	min := int(math.Inf(1))

	for i := 1; i < len(nodes); i++ {
		diff := nodes[i] - nodes[i-1]
		if diff < 0 {
			diff *= -1
		}
		if diff < min {
			min = diff
		}
	}

	return min
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
