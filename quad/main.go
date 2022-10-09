/**
 * Definition for a QuadTree node.
 * type Node struct {
 *     Val bool
 *     IsLeaf bool
 *     TopLeft *Node
 *     TopRight *Node
 *     BottomLeft *Node
 *     BottomRight *Node
 * }
 */
package main

func main() {}

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	// each time divide by 4 and of the 4 recurse
	if len(grid) == 0 {
		return nil
	}
	return process(grid, 0, len(grid), 0, len(grid))
}

func process(grid [][]int, r1, r2, c1, c2 int) *Node {

	var cnt0, cnt1 = 0, 0

	for i := r1; i < r2; i++ {
		for j := c1; j < c2; j++ {
			if grid[i][j] == 0 {
				cnt0++
			} else {
				cnt1++
			}
		}
	}

	if cnt0 == 0 {
		return &Node{
			Val:    true,
			IsLeaf: true,
		}
	}
	if cnt1 == 0 {
		return &Node{
			Val:    false,
			IsLeaf: true,
		}
	}
	return &Node{
		Val:         false,
		IsLeaf:      false,
		TopLeft:     process(grid, r1, (r1+r2)/2, c1, (c1+c2)/2),
		TopRight:    process(grid, r1, (r1+r2)/2, (c1+c2)/2, c2),
		BottomLeft:  process(grid, (r1+r2)/2, r2, c1, (c1+c2)/2),
		BottomRight: process(grid, (r1+r2)/2, r2, (c1+c2)/2, c2),
	}
}

func dfs(grid [][]int, x, y, l int) *Node {
	if l == 1 {
		var value bool
		if grid[x][y] == 1 {
			value = true
		}
		return &Node{
			Val:    value,
			IsLeaf: false,
		}
	}
	half := l / 2
	topLeft := dfs(grid, x, y, half)
	topRight := dfs(grid, x, y+half, half)
	bottomLeft := dfs(grid, x+half, y, half)
	bottomRight := dfs(grid, x+half, y+half, half)
	if topLeft.IsLeaf && topRight.IsLeaf && bottomLeft.IsLeaf && bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val == bottomLeft.Val == bottomRight.Val {
		return &Node{
			Val:    topLeft.Val,
			IsLeaf: true,
		}
	}
	return &Node{true, false, topLeft, topRight, bottomLeft, bottomRight}
}

// class Solution:
//     def construct(self, grid: List[List[int]]) -> 'Node':
//         def process(r1, r2, c1, c2):
//             cnt0 = 0
//             cnt1 = 0
//             for i in range(r1, r2):
//                 for j in range(c1, c2):
//                     if grid[i][j] == 0:
//                         cnt0 += 1
//                     else:
//                         cnt1 += 1
//             if cnt0 == 0:
//                 node = Node(1, True)
//                 return node
//             if cnt1 == 0:
//                 node = Node(0, True)
//                 return node
//             node = Node(0, False)
//             node.topLeft = process(r1, (r1 + r2)//2, c1, (c1+c2)//2)
//             node.topRight = process(r1, (r1 + r2)//2, (c1+c2)//2, c2)
//             node.bottomLeft = process((r1 + r2)//2, r2, c1, (c1+c2)//2)
//             node.bottomRight = process((r1 + r2)//2, r2, (c1+c2)//2, c2)
//             return node
//         if len(grid) == 0:
//             return None
//         return process(0, len(grid), 0, len(grid))

// # Time:  O(n)
// # Space: O(h)

// class Node(object):
//     def __init__(self, val, isLeaf, topLeft, topRight, bottomLeft, bottomRight):
//         self.val = val
//         self.isLeaf = isLeaf
//         self.topLeft = topLeft
//         self.topRight = topRight
//         self.bottomLeft = bottomLeft
//         self.bottomRight = bottomRight

// class Solution(object):
//     def construct(self, grid):
//         """
//         :type grid: List[List[int]]
//         :rtype: Node
//         """
//         def dfs(grid, x, y, l):
//             if l == 1:
//                 return Node(grid[x][y] == 1, True, None, None, None, None)
//             half = l // 2
//             topLeftNode = dfs(grid, x, y, half)
//             topRightNode = dfs(grid, x, y+half, half)
//             bottomLeftNode = dfs(grid, x+half, y, half)
//             bottomRightNode = dfs(grid, x+half, y+half, half)
//             if topLeftNode.isLeaf and topRightNode.isLeaf and \
//                bottomLeftNode.isLeaf and bottomRightNode.isLeaf and \
//                topLeftNode.val == topRightNode.val == bottomLeftNode.val == bottomRightNode.val:
//                 return Node(topLeftNode.val, True, None, None, None, None)
//             return Node(True, False, topLeftNode, topRightNode, bottomLeftNode, bottomRightNode)

//         if not grid:
//             return None
//         return dfs(grid, 0, 0, len(grid))
