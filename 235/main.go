package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

// 	if root == nil {
// 		return root
// 	}
// 	arr := []int{root.Val}
// 	if root.Left != nil {
// 		arr = append(arr, root.Left.Val)
// 	}
// 	if root.Right != nil {
// 		arr = append(arr, root.Right.Val)
// 	}
// 	if inArray(arr, p.Val) && inArray(arr, q.Val) {
// 		return root
// 	}
// 	left := lowestCommonAncestor(root.Left, p, q)
// 	if left != nil {
// 		return left
// 	}
// 	right := lowestCommonAncestor(root.Right, p, q)
// 	if right != nil {
// 		return right
// 	}
// 	return nil
// }

// func inArray(arr []int, target int) bool {

// 	for _, n := range arr {
// 		if n == target {
// 			return true
// 		}
// 	}
// 	return false
// }

// traverse the tree using preorder = root, left, right
// check the value of left and right tree
// - equal to p and q then stop - what can I do to check this
// - I can do check in array
// - if not equal, traverse left first - when left is not nil - then return
// - otherwise traverse right as a return value
// when to stop the recursion - only stop when value is nil so we can do for checking

// use preorder tree traverse
// when there is still right and left tree to traverse we will do this
// and when we reach the point that we cannot traverse further
// check the val - like when something is a leaf
// 2 - 0 - 4 -> looking for value of 0 and 4

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	rootVal := root.Val
	pVal := p.Val
	qVal := q.Val

	if pVal > rootVal && qVal > rootVal {
		return lowestCommonAncestor(root.Right, p, q)
	}
	if pVal < rootVal && qVal < rootVal {
		return lowestCommonAncestor(root.Left, p, q)
	}
	return root
}

// looking for 0 and 5
// doing the traversal -> 2 - 0 - 4 - 3 - 5
// as I traverse further to left and right
// when I hit the left -> 0, 0, 5 - I will check if everything in there if not then I will return nil
// when I hit the right ->

// we need to use BST properties here
// value of left smaller than root and right larger than root
