package main

import "fmt"

func main() {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	// root.Next.Next.Next = &ListNode{Val: 3}
	// root.Next.Next.Next.Next = &ListNode{Val: 4}
	// root.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	// root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	new := removeElementsIter(root, 2)

	fmt.Printf("%#v\n", new.Val)
	fmt.Printf("%#v\n", new.Next.Val)
	// fmt.Printf("%#v\n", new.Next.Next.Val)
	// fmt.Printf("%#v\n", new.Next.Next.Next.Val)
	// fmt.Printf("%#v\n", new.Next.Next.Next.Next)
	// fmt.Printf("%#v\n", new.Next.Next.Next.Next.Val)
	// fmt.Printf("%#v\n", new.Next.Next.Next.Next.Next)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElementsIter(head *ListNode, val int) *ListNode {

	var sentinel = &ListNode{}
	sentinel.Next = head

	prev, current := sentinel, head

	for current != nil {
		if current.Val == val {
			prev.Next = current.Next
		} else {
			prev = current
		}
		fmt.Printf("%#v\n", prev)
		current = current.Next
	}

	return sentinel.Next
}

func removeElementsRecur(head *ListNode, val int) *ListNode {

	// base case
	if head == nil {
		return nil
	}

	// include this
	if head.Val != val {
		head.Next = removeElementsRecur(head.Next, val)
		// dont include this and move on to next
	} else {
		head = removeElementsRecur(head.Next, val)
	}

	return head
}
