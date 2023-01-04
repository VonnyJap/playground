package main

import "fmt"

func main() {

	root := &TreeNode{Val: 10}
	root.Left = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 7}
	root.Right = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 18}

	// 10,5,15,3,7,null,18
	fmt.Println(rangeSumBST(root, 7, 15))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// since this is BST - we can do inorder so it will be sorted already
// and then we can add everything within the range
func rangeSumBST(root *TreeNode, low int, high int) int {

	nodes := inorderTraversal(root)

	sum := 0
	for _, n := range nodes {
		if n >= low && n <= high {
			sum += n
		}
	}

	return sum
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
