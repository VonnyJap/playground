package main

import "fmt"

func main() {
	fmt.Printf("%+v\n", allPossibleFBT(8))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var memo = map[int][]*TreeNode{
	0: []*TreeNode{},
	1: []*TreeNode{&TreeNode{}},
}

func allPossibleFBT(n int) []*TreeNode {

	if r, ok := memo[n]; ok {
		return r
	}
	var results = []*TreeNode{}
	for i := 0; i < n; i++ {
		left := allPossibleFBT(i)
		right := allPossibleFBT(n - i - 1)
		for _, l := range left {
			for _, r := range right {
				bns := &TreeNode{}
				bns.Left = l
				bns.Right = r
				results = append(results, bns)
			}
		}
	}
	memo[n] = results
	return memo[n]
}
