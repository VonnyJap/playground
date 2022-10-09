package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	printList(middleNode(head))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func middleNode(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	ptr1 := head
	ptr2 := head

	for ptr2 != nil {
		ptr2 = ptr2.Next
		if ptr2 == nil {
			break
		}
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1
}
