package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// this is a binary tree hence do traversal
// 5 - 6 - 2 - 7 - 4 -> looking for 6 and 4
// in order for the function not to return earlier
// the function needs to keep on recursing until not nil

// at 5
// traverse 6 -> found 6
// traverse 2 -> found 2 - 7 - 4
// when p and q contained in that child -> traverse further
// until it will eventually stop
// things need to happen like this to stop
// when root.Val equal to p or q - then either one child need to contain the other one
// else the two children will have either one and then we stop and return root

// how can make this thing faster

// the algo is
// traverse the tree and set the parent for each node until they are found
// and then use another hashmap to store all values of all parents of p
// and then check q and return as soon as we find the value of common parents
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	stack := []*TreeNode{root}
	parent := map[*TreeNode]*TreeNode{root: nil}

	for len(stack) > 0 {
		if _, ok := parent[p]; ok {
			if _, ok := parent[q]; ok {
				break
			}
		}
		current := stack[0]
		stack = stack[1:]
		if current.Left != nil {
			parent[current.Left] = current
			stack = append(stack, current.Left)
		}
		if current.Right != nil {
			parent[current.Right] = current
			stack = append(stack, current.Right)
		}
	}

	parentP := map[*TreeNode]bool{}

	for p != nil {
		parentP[p] = true
		p = parent[p]
	}

	var common *TreeNode
	for q != nil {
		if _, ok := parentP[q]; ok {
			common = q
			break
		}
		q = parent[q]
	}
	return common
}

// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

// 	if root == nil {
// 		return root
// 	}
// 	leftVals := preorderTraversal(root.Left)
// 	if inArray(leftVals, p.Val) && inArray(leftVals, q.Val) {
// 		return lowestCommonAncestor(root.Left, p, q)
// 	}
// 	rightVals := preorderTraversal(root.Right)
// 	if inArray(rightVals, p.Val) && inArray(rightVals, q.Val) {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	}
// 	return root
// }

// func preorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}

// 	if root.Left == nil && root.Right == nil {
// 		return []int{root.Val}
// 	}

// 	var nodes []int
// 	nodes = append(nodes, []int{root.Val}...)
// 	if root.Left != nil {
// 		nodes = append(nodes, preorderTraversal(root.Left)...)
// 	}
// 	if root.Right != nil {
// 		nodes = append(nodes, preorderTraversal(root.Right)...)
// 	}
// 	return nodes
// }

// func inArray(arr []int, target int) bool {

// 	for _, n := range arr {
// 		if n == target {
// 			return true
// 		}
// 	}
// 	return false
// }
