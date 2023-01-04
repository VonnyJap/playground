package main

import (
	"fmt"
	"math"
)

func main() {
	root := &TreeNode{Val: 1}
	fmt.Println(closestValue(root, 4.428571))

	root = &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 5}

	fmt.Println(closestValue(root, 3.714286))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {

	values := inorderTraversal(root)
	// fmt.Println(values)
	closest := values[0]
	diff := math.Abs(float64(values[0]) - target)

	for _, v := range values[1:] {
		tempDiff := math.Abs(float64(v) - target)
		if tempDiff > diff {
			break
		}
		diff = tempDiff
		closest = v
	}

	return closest
}

// 1,2,3,4,5 - 3.714286

// 1 - 3.714286 = 2.714286
// 2 - 3.714286 = 1.714286
// 3 - 3.714286 = 0.714286
// 4 - 3.714286 = 0.285714

// again this is a tree so we need to traverse to check the value
// 1. do inorder traverse to get all values from the tree
// 2. the inorder traverse should be giving value from small to big
// 3. use modified binary search to get the closest value?

// lets do linear search and we keep track the absolute difference
// as soon as we find an absolute difference from small to big then we break the loop

func inorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	result := append(inorderTraversal(root.Left), []int{root.Val}...)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}
