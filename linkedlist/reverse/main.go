package main

import "fmt"

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}

	rvr := reverse(head)
	printList(rvr)
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

func reverse(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var reverseList *ListNode = nil

	for head != nil {
		newNext := reverseList
		currNext := head.Next
		reverseList = head
		reverseList.Next = newNext
		head = currNext
	}

	return reverseList
}
