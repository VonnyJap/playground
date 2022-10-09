package main

import "fmt"

func main() {

	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 4}
	fmt.Println(getDirections(root, 3, 6))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lca(root *TreeNode, startValue int, destValue int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == startValue {
		return root
	}
	if root.Val == destValue {
		return root
	}
	left := lca(root.Left, startValue, destValue)
	right := lca(root.Right, startValue, destValue)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func getPath(root *TreeNode, value int, path *[]string) bool {
	if root == nil {
		return false
	}
	if root.Val == value {
		return true
	}

	*path = append(*path, "L")
	resp := getPath(root.Left, value, path)
	if resp {
		return true
	}

	*path = (*path)[0 : len(*path)-1]
	*path = append(*path, "R")
	resp = getPath(root.Right, value, path)
	if resp {
		return true
	}

	*path = (*path)[0 : len(*path)-1]
	return false
}

func getDirections(root *TreeNode, startValue int, destValue int) string {

	root = lca(root, startValue, destValue)

	var p1 = []string{}
	var p2 = []string{}

	getPath(root, startValue, &p1)
	getPath(root, destValue, &p2)

	fmt.Println("p1: ", p1)
	fmt.Println("p2: ", p2)

	for i := range p1 {
		p1[i] = "U"
	}

	dir := ""
	for _, v := range p1 {
		dir += v
	}
	for _, v := range p2 {
		dir += v
	}
	return dir
}
