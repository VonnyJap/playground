package main

import "sort"

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {

	nodes := preorderTraversal(root)

	// remove dups
	nodesMap := map[int]bool{}

	for _, n := range nodes {
		if _, ok := nodesMap[n]; ok {
			continue
		}
		nodesMap[n] = true
	}

	nodes = []int{}
	for n := range nodesMap {
		nodes = append(nodes, n)
	}

	sort.Ints(nodes)

	if len(nodes) == 0 {
		return 0
	}
	if len(nodes) == 1 {
		return nodes[0]
	}

	return nodes[1]
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	var nodes []int
	nodes = append(nodes, []int{root.Val}...)
	if root.Left != nil {
		nodes = append(nodes, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nodes = append(nodes, preorderTraversal(root.Right)...)
	}
	return nodes
}
