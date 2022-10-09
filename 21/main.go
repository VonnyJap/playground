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

	var mergedList = &ListNode{}
	var curent = mergedList

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curent.Next = list1
			list1 = list1.Next
		} else {
			curent.Next = list2
			list2 = list2.Next
		}
		curent = curent.Next
	}

	for list1 != nil {
		curent.Next = list1
		list1 = list1.Next
		curent = curent.Next
	}

	for list2 != nil {
		curent.Next = list2
		list2 = list2.Next
		curent = curent.Next
	}

	return mergedList.Next
}

// the list is sorted
// two pointers
// 1: point to the head of linked list 1
// 2: point to the head of linked list 2
// for loop - will run until it finds either one of the linked lists have no more Next node
// and then return
