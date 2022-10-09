package main

import "fmt"

func main() {
	start := &ListNode{Val: 1}
	start.Next = &ListNode{Val: 1}
	start.Next.Next = &ListNode{Val: 2}
	start.Next.Next.Next = &ListNode{Val: 3}
	start.Next.Next.Next.Next = &ListNode{Val: 3}

	new := deleteDuplicates(start)
	fmt.Printf("%#v\n", new.Val)
	fmt.Printf("%#v\n", new.Next.Val)
	fmt.Printf("%#v\n", new.Next.Next.Val)
	fmt.Printf("%#v\n", new.Next.Next.Next)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// this is a sorted linked list
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	if head.Val == head.Next.Val {
		head = head.Next
		head = deleteDuplicates(head)
	} else {
		head.Next = deleteDuplicates(head.Next)
	}
	return head
}
