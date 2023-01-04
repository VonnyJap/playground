package main

import "fmt"

func main() {

	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 1}
	root.Left.Left = &TreeNode{Val: 6}
	root.Left.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 8}
	root.Left.Right.Left = &TreeNode{Val: 7}
	root.Left.Right.Right = &TreeNode{Val: 4}

	fmt.Println(getLeaf(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	root1Leaf := getLeaf(root1)
	root2Leaf := getLeaf(root2)

	if len(root1Leaf) != len(root2Leaf) {
		return false
	}

	for i := range root1Leaf {
		if root1Leaf[i] != root2Leaf[i] {
			return false
		}
	}

	return true
}

// do traversal first but only add if it is a leaf

func getLeaf(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}

	if root.Right == nil && root.Left == nil {
		return []int{root.Val}
	}

	return append(getLeaf(root.Left), getLeaf(root.Right)...)
}
