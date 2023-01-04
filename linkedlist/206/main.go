package main

import "fmt"

func main() {
	node := &ListNode{Val: 1}
	node.Next = &ListNode{Val: 2}
	node.Next.Next = &ListNode{Val: 3}

	// result := reverseList(node)

	result := reverseListInPlace(node)
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

func reverseList(head *ListNode) *ListNode {

	var result *ListNode

	for head != nil {
		prev := result
		next := head.Next
		result = head
		result.Next = prev
		head = next
	}

	return result
}

// 1 - 2 - 3
// 1. prev = nil, result = 1, next = nil
// 2. prev = (1,nil), result = 2, next =  (1,nil)

func reverseListInPlace(head *ListNode) *ListNode {
	curr := head
	fmt.Println("curr:", curr)
	if curr == nil || curr.Next == nil {
		return curr
	}
	ptr := reverseListInPlace(curr.Next)
	fmt.Println("ptr:", ptr)
	curr.Next.Next = curr
	curr.Next = nil
	fmt.Println("new ptr:", ptr)
	fmt.Println("new curr:", curr)
	return ptr
}
