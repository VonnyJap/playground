package main

func main() {}

// because this is a tree - traversal or dfs
// when traversing - need to decide which way to go

// 3 -> 9
// 3 -> 20 -> 15 -> 7

// since I need to know the length of children and decide which one is longer - recur ove the children to get the depth
// then I need to do postprocessing of the current node

// what is the base condition then?
// when node is nil then - we want to return 0?
// when node does not have children we want to return 1?

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		return 1 + leftDepth
	}
	return rightDepth
}
