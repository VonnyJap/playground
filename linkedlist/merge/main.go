package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 3}
	fmt.Println("print l1 before")
	printList(l1)

	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 5}
	l2.Next.Next = &ListNode{Val: 6}
	fmt.Println("print l2 before")
	printList(l2)

	merge(l1, l2)

	fmt.Println("print l1 after")
	printList(l1)

	fmt.Println("print l2 after")
	printList(l2)

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

// test with 1 and 1
func merge(l1, l2 *ListNode) {

	var temp_1 = l1

	for temp_1 != nil && l2 != nil {

		temp_1_next := temp_1.Next
		l2_next := l2.Next

		l2.Next = temp_1_next
		temp_1.Next = l2

		l2 = l2_next
		fmt.Printf("%#v\n", l2)
		temp_1 = temp_1_next
	}

	fmt.Printf("%#v\n", l2)

	return
}
