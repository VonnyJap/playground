package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next.Next = &ListNode{Val: 5}
	l1.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	printList(GroupReverse(l1, 3))

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

// how to change a pointer?
func GroupReverse(list *ListNode, group int) *ListNode {

	// each time cut - reverse and then connect to the result something like that

	var result = &ListNode{}
	var temp = result

	for list != nil {
		temp.Next = Reverse(Cut(list, group))
		i := 0
		for i < group && list != nil {
			list = list.Next
			temp = temp.Next
			i++
		}
	}
	return result.Next
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

func Cut(list *ListNode, count int) *ListNode {

	values := []int{}
	var result = &ListNode{}
	var temp = result
	var curr = list
	i := 0
	// printList(list)
	for curr != nil && i < count {
		values = append(values, curr.Val)
		curr = curr.Next
		i++
	}
	for _, n := range values {
		temp.Next = &ListNode{Val: n}
		temp = temp.Next
	}
	return result.Next
}
