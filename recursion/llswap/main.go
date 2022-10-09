package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}

	printList(head)

	printList(llswap(head))
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

func llswap(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var current *ListNode = head
	var next *ListNode = head.Next
	var next_next *ListNode = head.Next.Next
	head = next
	head.Next = current
	head.Next.Next = llswap(next_next)

	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
