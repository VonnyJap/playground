package main

import "fmt"

func main() {

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	// head.Next.Next.Next.Next = &ListNode{Val: 5}

	node := reverseList(head)
	fmt.Printf("%#v\n", node.Val)
	fmt.Printf("%#v\n", node.Next.Val)
	fmt.Printf("%#v\n", node.Next.Next.Val)
	fmt.Printf("%#v\n", node.Next.Next.Next)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var stack = []int{}

	// why this works? because it is liked we keep track of the address of this
	// hence we are not really affected when this get changed
	var newHead *ListNode = head

	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	// fmt.Printf("boom: %#v\n", newHead.Val)
	// fmt.Printf("boom: %#v\n", newHead.Next.Val)
	// fmt.Printf("boom: %#v\n", newHead.Next.Next.Val)
	// fmt.Printf("boom: %#v\n", newHead.Next.Next.Next)

	var yoHead *ListNode = newHead

	for len(stack) > 0 {
		// fmt.Printf("%#v\n", newHead)

		newHead.Val = stack[len(stack)-1]
		newHead = newHead.Next
		stack = stack[:len(stack)-1]
	}
	return yoHead
}
