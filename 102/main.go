package main

import "fmt"

func main() {

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 8}
	root.Right.Left.Left = &TreeNode{Val: 12}

	fmt.Println(levelOrder(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// level to level - bst
// postorder traverse

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	left := levelOrder(root.Left)
	right := levelOrder(root.Right)
	length := len(left)
	if len(left) < len(right) {
		length = len(right)
	}
	var result = make([][]int, length)
	for i, l := range left {
		result[i] = append(result[i], l...)
	}
	for i, r := range right {
		result[i] = append(result[i], r...)
	}
	return append([][]int{{root.Val}}, result...)
}

// [3]
// [9, 20]

// [15, 7]

// based on the observation - this is post order traverse
// we want to make sure that we join the level together
// for each of the left and right we need to combine those as well
// by looping it through

// 1. base case
// 	- root == nil
// 	- root.Left && root.Right == nil
// 2. process the left and right
// 3. combine the result between these twos
// 4. add root in front of them
