package main

import "fmt"

func main() {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 0}
	root.Next.Next.Next = &ListNode{
		Val:  4,
		Next: root.Next,
	}
	fmt.Println(hasCycle(root))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// var a []*ListNode

// func hasCycle(head *ListNode) bool {
// 	if head == nil {
// 		return false
// 	}
// 	if head.Next == nil {
// 		return false
// 	}
// 	for _, v := range a {
// 		if v.Val == head.Val && v.Next.Val == head.Next.Val {
// 			return true
// 		}
// 	}
// 	a = append(a, head)
// 	return hasCycle(head.Next)
// }

// for tree or link sometimes we can just use iterative
