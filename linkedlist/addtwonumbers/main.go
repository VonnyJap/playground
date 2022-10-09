package main

import "fmt"

func main() {

	l1 := &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 2}

	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 5}
	// l2.Next.Next = &ListNode{Val: 1}
	// [9,9,9,9,9,9,9]
	// [9,9,9,9]
	result := addTwoNumbers(l1, l2)
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	n1 := getSize(l1)
	n2 := getSize(l2)
	diff := n2 - n1
	if diff == 0 {
		return add(l1, l2, 0)
	}
	if diff < 0 {
		l1, l2 = l2, l1
		diff = diff * (-1)
	}
	var temp = l1
	for temp.Next != nil {
		temp = temp.Next
	}
	i := 0
	for i < diff {
		temp.Next = &ListNode{Val: 0}
		temp = temp.Next
		i++
	}
	printList(l1)
	return add(l1, l2, 0)
}

func getSize(l *ListNode) int {
	n := 0
	for l != nil {
		n++
		l = l.Next
	}
	return n
}

func add(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// l1 and l2 both same
	if l1 == nil {
		if carry == 0 {
			return nil
		}
		return &ListNode{Val: carry}
	}
	var result = &ListNode{}
	sum := l1.Val + l2.Val + carry
	result.Val = sum % 10
	carry = int(sum / 10)
	result.Next = add(l1.Next, l2.Next, carry)
	return result
}

// get the size and if not same size then we will add 0 to the back
// and after that we will add one by one
