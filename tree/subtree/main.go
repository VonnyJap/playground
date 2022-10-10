package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// root vs subRoot
// traversing root: 3 - 4 - 1 - 2 - 5
// traversing subRoot: 4 - 1 - 2
// there are two cases to think about here
// 1. when root and subRoot not same value yet - traverse only the root - comparing root.Left and root.Right with subRoot
// 2. when value the same then - compare root.Left with subRoot.Left and root.Right with subRoot.Right

// -> if both nil -> return true
// -> if one of them nil -> return false
// root.Val == subRoot.Val -> isSubtree(root.Left, subRoot.Left) && isSubtree(root.Right, subRoot.Right)
// return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

	if root == nil && subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if subRoot == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
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
