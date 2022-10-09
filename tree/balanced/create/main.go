package main

import "fmt"

func main() {
	node := sortedArrayToBST([]int{-10, -3, 0, 5, 9})

	fmt.Printf("%#v\n", node.Val)
	fmt.Printf("%#v\n", node.Left.Val)
	fmt.Printf("%#v\n", node.Left.Left.Val)
	// fmt.Printf("%#v\n", node.Left.Left.Left.Val)

	fmt.Printf("%#v\n", node.Right.Val)
	fmt.Printf("%#v\n", node.Right.Left.Val)
	fmt.Printf("%#v\n", node.Right.Left.Right)

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	if len(nums) == 2 {
		return &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val: nums[0],
			},
		}
	}
	rootIndex := int(len(nums) / 2)
	return &TreeNode{
		Val:   nums[rootIndex],
		Left:  sortedArrayToBST(nums[0:rootIndex]),
		Right: sortedArrayToBST(nums[rootIndex+1:]),
	}
}
