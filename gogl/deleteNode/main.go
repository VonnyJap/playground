package main

import "fmt"

func main() {

	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 9}

	current := deleteNode(head, 1)
	printList(current)
	// 4,5,1,9 - 5
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, key int) *ListNode {

	current := head
	var prev *ListNode

	for current != nil {
		if current.Val == key {
			if prev == nil {
				return current.Next
			}
			prev.Next = current.Next
			break
		}
		prev = current
		current = current.Next
	}

	return head
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

// loop until next is pointing to nil
// we need to keep track of prev as well
// prev.next = current.next
