package main

import "fmt"

func main() {

	n := &ListNode{Val: 1}
	n.Next = &ListNode{Val: 2}
	// n.Next.Next = &ListNode{Val: 3}
	// n.Next.Next.Next = &ListNode{Val: 4}
	// n.Next.Next.Next.Next = &ListNode{Val: 5}

	printList(removeNthFromEnd(n, 1))

}

// reverse and then reverse again

type ListNode struct {
	Val  int
	Next *ListNode
}

func Reverse(list *ListNode) *ListNode {

	var new *ListNode = nil
	for list != nil {
		new_next := new
		list_next := list.Next
		new = list
		new.Next = new_next
		list = list_next
	}

	return new
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	var reverse *ListNode = Reverse(head)

	if n == 1 {
		return Reverse(reverse.Next)
	}
	var curr *ListNode = reverse
	var prev *ListNode = reverse

	i := 0
	for curr != nil && i < n {
		if i == n-1 {
			// remove current and attach the current.Next as prev.Next
			prev.Next = curr.Next
		}
		prev = curr
		curr = curr.Next
		i++
	}

	return Reverse(reverse)
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
