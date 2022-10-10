package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Most of the time if given binary tree what can be done is to traverse the tree
// example 1: 1 - 2 - 3 and 1 - 2 - 3
// example 2: 1 - 2 and 1 - null - 2
// and then i decide to pre, post, or between
// in this case I am going to check the pre which is the current node
// if current not same -> return false
// check if left and left same -> if not then false
// check if right and right same -> if not then false
// base case to determine when I will stop
// if both nil then I am stop - return true
// else I am returning false

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
