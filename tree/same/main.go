package main

import "fmt"

func main() {

	// one := TreeNode{Val: 1}
	two := TreeNode{Val: 2}

	a := TreeNode{
		Val:  3,
		Left: &two,
	}
	b := TreeNode{
		Val:   3,
		Right: &two,
	}
	fmt.Println(isSameTree(&a, &b))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// this needs to be recursive as well
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
