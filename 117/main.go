package main

func main() {}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// since this is looking for the neighbor
// to begin how to recurse and when to stop?

func connect(root *Node) *Node {

}

// Input: root = [1,2,3,4,5,null,7]
// Output: [1,#,2,3,#,4,5,7,#]

// val = 1, Next = nil
// val = 2, Next = 3
// val = 3, Next = nil
// val = 4, Next = 5
// val = 5, Next = 7
// val = 7, Next = nil

// how do I know the thing at the same level
