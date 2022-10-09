package main

import "fmt"

func main() {
	tree := &TreeNode{Val: 2}
	tree.Left = &TreeNode{Val: 1}
	tree.Right = &TreeNode{Val: 3}

	inverted := invertTree(tree)
	fmt.Println(inverted.Val)
	fmt.Printf("%+v\n", inverted.Left)
	fmt.Printf("%+v\n", inverted.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}
	right := root.Right
	left := root.Left
	root.Left = invertTree(right)
	root.Right = invertTree(left)

	return root
}

// so because this is a tree - I need to traverse the tree
// when I traverse tree - it is a recursion process
// and since this is a recursion - I will need to find a base case which happened when root == nil
// if root == nil -> then stop
// otherwise do some processing
// root.Left = invertTree(root.Right)
// root.Right = invertTree(root.Left)
// return root

// 2, 1, 3 -> 2, 3, 1

// 4, 2, 1, 3 -> 4, 2, 3, 1
