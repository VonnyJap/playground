package main

import "fmt"

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(averageOfLevels(root))

	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Left.Left = &TreeNode{Val: 15}
	root1.Left.Right = &TreeNode{Val: 7}

	fmt.Println(averageOfLevels(root1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {

	if root == nil {
		return []float64{}
	}

	type TreeLevel struct {
		Tree  *TreeNode
		Level int
	}

	queue := []*TreeLevel{{root, 0}}
	result := []float64{}

	current := 0
	count := 0
	prevLevel := 0

	for len(queue) > 0 {

		now := queue[0]

		queue = queue[1:]
		if now.Tree.Left != nil {
			queue = append(queue, &TreeLevel{now.Tree.Left, now.Level + 1})
		}
		if now.Tree.Right != nil {
			queue = append(queue, &TreeLevel{now.Tree.Right, now.Level + 1})
		}
		if now.Level == prevLevel {
			current += now.Tree.Val
			count++
		} else {
			result = append(result, float64(current)/float64(count))
			current = now.Tree.Val
			count = 1
		}
		prevLevel = now.Level
	}

	if count > 0 {
		result = append(result, float64(current)/float64(count))
	}

	return result
}
