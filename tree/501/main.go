package main

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	fmt.Println(findMode(root))

	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 2}

	fmt.Println(findMode(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [1,2,2] -
func findMode(root *TreeNode) []int {

	nodes := inorderTraversal(root)
	fmt.Println(nodes)
	max := 1
	tempMax := 1
	result := []int{nodes[0]}

	for i := 1; i < len(nodes); i++ {
		if nodes[i] == nodes[i-1] {
			tempMax++
		} else {
			tempMax = 1
		}
		if tempMax > max {
			result = []int{nodes[i]}
			max = tempMax
		} else if tempMax == max {
			result = append(result, nodes[i])
		}
		fmt.Println(max, tempMax, result)
	}

	return result

}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if (root.Left == nil) && (root.Right == nil) {
		return []int{root.Val}
	}
	var nodes []int
	nodes = append(nodes, inorderTraversal(root.Left)...)
	nodes = append(nodes, []int{root.Val}...)
	nodes = append(nodes, inorderTraversal(root.Right)...)
	return nodes
}

// we can traverse the node from left root right and we should have everything in order
// and then we can use the max and value
