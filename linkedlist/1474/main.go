package main

import "fmt"

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 11}

	result := deleteNodes(head, 1, 3)
	printList(result)
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

func deleteNodes(head *ListNode, m int, n int) *ListNode {

	current := head
	var prev *ListNode = nil
	i := 0

	for current != nil {
		if i >= m && i < m+n {
			prev.Next = current.Next
		} else {
			prev = current
		}
		current = current.Next
		i++
		if i == m+n {
			i = 0
		}
		// we need to somehow change the value of i
		// if it reaches here - it will become n
	}

	return head
}
