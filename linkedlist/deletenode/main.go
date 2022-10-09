package main

import "fmt"

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	new := head.deleteNode(1)
	fmt.Printf("new val: %#v\n", new)
	// fmt.Printf("new next val: %d\n", new.Next.Val)
	// fmt.Printf("new next val: %#v\n", new.Next.Next)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// we need to handle the case where the node that we are trying to delete is the head

func (node *ListNode) deleteNode(data int) *ListNode {
	var temp = node
	var prev = node

	for {
		if temp == nil {
			break
		}
		if temp.Val == data {
			// case of deleting the head
			if temp == prev {
				if temp.Next != nil {
					return temp.Next
				}
				return nil
			}
			prev.Next = temp.Next
			break
		}
		prev = temp
		temp = temp.Next
	}

	return node
}
