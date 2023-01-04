package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}

	root.Right = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(findTilt(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// can do better with memoization?
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	sumLeft := 0
	for _, l := range left {
		sumLeft += l
	}

	sumRight := 0
	for _, r := range right {
		sumRight += r
	}

	return int(math.Abs(float64(sumLeft)-float64(sumRight))) + findTilt(root.Left) + findTilt(root.Right)
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

// need to get the tilt and then add it up

// this problem statement is to find tilt - (left - right) and then add the tilt from left and also right as well
// and then sum it up

// steps:
// 1. get the sum of tilt from left, sum of tilt from right as well
// 2. get the current tilt - substract right and left
// 3. sum all of them together
// 4. base case
// 	- when the tree is nil
// 	- when the left and right is nil

// another way I can do is: inorder traverse and then I find where the root is and then find the diff
