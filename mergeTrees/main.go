package main

import "fmt"

func main() {

	tree1 := &TreeNode{Val: 1}
	tree1.Left = &TreeNode{Val: 3}
	tree1.Right = &TreeNode{Val: 2}
	tree1.Left.Left = &TreeNode{Val: 5}

	tree2 := &TreeNode{Val: 2}
	tree2.Left = &TreeNode{Val: 1}
	tree2.Right = &TreeNode{Val: 3}
	tree2.Left.Right = &TreeNode{Val: 4}
	tree2.Right.Right = &TreeNode{Val: 7}

	tree3 := mergeTrees(tree1, tree2)

	fmt.Println(tree3.Val)
	fmt.Println(tree3.Left.Val)
	fmt.Println(tree3.Right.Val)
	fmt.Println(tree3.Left.Left.Val)
	fmt.Println(tree3.Left.Right.Val)
	fmt.Println(tree3.Right.Left)
	fmt.Println(tree3.Right.Right.Val)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// what will be the base condition?
// basically we need to check

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	if root2 == nil {
		return root1
	}
	if root1 == nil {
		return root2
	}
	if root1.Left == nil && root1.Right == nil && root2.Left == nil && root2.Right == nil {
		return &TreeNode{Val: root1.Val + root2.Val}
	}

	tree := &TreeNode{}
	val := 0
	if root1 != nil {
		val += root1.Val
	}
	if root2 != nil {
		val += root2.Val
	}
	tree.Val = val
	tree.Left = mergeTrees(root1.Left, root2.Left)
	tree.Right = mergeTrees(root1.Right, root2.Right)

	return tree
}
