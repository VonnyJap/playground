package main

import "fmt"

func main() {
	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}
	printList(mergeTwoLists(list1, list2))
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	ptr := head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			ptr.Next = list1
			list1 = list1.Next
		} else if list1.Val > list2.Val {
			ptr.Next = list2
			list2 = list2.Next
		}
		ptr = ptr.Next
	}

	for list1 != nil {
		ptr.Next = list1
		list1 = list1.Next
		ptr = ptr.Next
	}

	for list2 != nil {
		ptr.Next = list2
		list2 = list2.Next
		ptr = ptr.Next
	}

	return head.Next
}
