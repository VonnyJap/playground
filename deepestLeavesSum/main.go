package main

import "fmt"

func main() {

	// Input: root = [1,2,3,4,5,null,6,7,null,null,null,null,8]

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Left.Left.Left = &TreeNode{Val: 7}
	root.Right.Right.Right = &TreeNode{Val: 8}

	fmt.Println(deepestLeavesSum(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// we can only know deepest by level
// we only know the level by doing BST

// how do I then tackle this problem?
// but this is a tree - we can still use the queue
// we need to keep a structure of level and the tree and put into the queue
// and we can define a map to store the level and components
// steps:
// 1. iterate over the tree and append it to the queue
// 2. exhaust the queue no more
// 3. each iter put into the map - level is the key and the value will be []int

type TreeNodeWithLevel struct {
	Item  *TreeNode
	Level int
}

func deepestLeavesSum(root *TreeNode) int {

	queue := []TreeNodeWithLevel{}
	queue = append(queue, TreeNodeWithLevel{root, 0})
	group := map[int][]int{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if _, ok := group[current.Level]; !ok {
			group[current.Level] = []int{current.Item.Val}
		} else {
			group[current.Level] = append(group[current.Level], current.Item.Val)
		}
		if current.Item.Left != nil {
			queue = append(queue, TreeNodeWithLevel{current.Item.Left, current.Level + 1})
		}
		if current.Item.Right != nil {
			queue = append(queue, TreeNodeWithLevel{current.Item.Right, current.Level + 1})
		}
	}

	deepest := len(group) - 1
	sum := 0
	for _, n := range group[deepest] {
		sum += n
	}

	return sum
}
