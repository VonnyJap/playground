package main

import (
	"fmt"
	"sort"
)

func main() {

	list1 := &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 4}
	list1.Next.Next = &ListNode{Val: 5}

	list2 := &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	list3 := &ListNode{Val: 2}
	list3.Next = &ListNode{Val: 6}

	merge := mergeKLists([]*ListNode{list1, list2, list3})
	printList(merge)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	result := []int{}

	// so each item here is the ListNode
	for i := 0; i < len(lists); i++ {
		a := lists[i]
		for {
			if a == nil {
				break
			}
			result = append(result, a.Val)
			a = a.Next
		}
	}

	sort.Ints(result)
	fmt.Println(result)

	var newHead *ListNode = &ListNode{}
	var yoHead = newHead

	for _, r := range result {
		yoHead.Next = &ListNode{Val: r}
		yoHead = yoHead.Next
	}

	return newHead.Next
}

func printList(n *ListNode) {
	// fmt.Printf("ListNode: %+v\n", n)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
