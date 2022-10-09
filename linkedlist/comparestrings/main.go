package main

import "fmt"

func main() {
	list1 := &Node{Val: "g"}
	list1.Next = &Node{Val: "e"}
	list1.Next.Next = &Node{Val: "e"}
	list1.Next.Next.Next = &Node{Val: "k"}
	list1.Next.Next.Next.Next = &Node{Val: "s"}
	list1.Next.Next.Next.Next.Next = &Node{Val: "a"}

	list2 := &Node{Val: "g"}
	list2.Next = &Node{Val: "e"}
	list2.Next.Next = &Node{Val: "e"}
	list2.Next.Next.Next = &Node{Val: "k"}
	list2.Next.Next.Next.Next = &Node{Val: "s"}
	list2.Next.Next.Next.Next.Next = &Node{Val: "b"}

	fmt.Println(compare(list1, list2))
}

type Node struct {
	Val  string
	Next *Node
}

func compare(n1, n2 *Node) int {

	// keep on looping
	// if same ok move on
	// if not then compare and return

	for n1 != nil && n2 != nil {
		if n1.Val > n2.Val {
			return 1
		}
		if n1.Val < n2.Val {
			return -1
		}
		n1 = n1.Next
		n2 = n2.Next
	}

	if n1 != nil {
		return 1
	}
	if n2 != nil {
		return -1
	}
	return 0
}
