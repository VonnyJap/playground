package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 20}
	tree.Left = &TreeNode{Val: 9}
	tree.Right = &TreeNode{Val: 25}
	tree.Left.Left = &TreeNode{Val: 5}
	tree.Left.Right = &TreeNode{Val: 12}
	tree.Left.Right.Left = &TreeNode{Val: 11}
	tree.Left.Right.Right = &TreeNode{Val: 14}

	fmt.Println(bstSuccessorSearch(tree, 9))
	fmt.Println(bstSuccessorSearch(tree, 11))
	fmt.Println(bstSuccessorSearch(tree, 14))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstSuccessorSearch(root *TreeNode, key int) int {

	// left
	// root
	// right

	elements := inOrderTraverse(root)
	i := 0
	for i < len(elements) {
		if elements[i] == key {
			break
		}
		i++
	}
	return elements[i+1]
}

func inOrderTraverse(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	result := []int{}
	result = append(result, inOrderTraverse(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inOrderTraverse(root.Right)...)

	return result

}
