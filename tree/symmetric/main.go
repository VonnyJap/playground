package main

import "fmt"

func main() {
	three := TreeNode{Val: 3}
	four := TreeNode{Val: 4}

	one := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &three,
			Right: &four,
		},
		Right: &TreeNode{
			Val:   2,
			Right: &three,
			Left:  &four,
		},
	}
	fmt.Println(isSymmetric(&one))

	c := TreeNode{Val: 3}
	b := TreeNode{
		Val:   2,
		Right: &c,
	}
	a := TreeNode{
		Val:   1,
		Left:  &b,
		Right: &b,
	}
	fmt.Println(isSymmetric(&a))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMirror(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 != nil && root2 != nil {
		return root1.Val == root2.Val && isMirror(root1.Left, root2.Right) && isMirror(root1.Right, root2.Left)
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}
