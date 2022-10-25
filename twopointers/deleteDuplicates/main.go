package main

import "fmt"

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val       int
	Duplicate bool
}

func deleteDuplicates(head *ListNode) *ListNode {

	ptr := head
	// stack := []Node{} // modify this later on to remove the duplicate and when that is done we will go through head again to get the right answer

	for ptr != nil {

	}

}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

// brute force
// we can loop until end
// each time go to stack and duplicate sign
// and then next if next next

// what about using two pointers
// 1, 1, 1, 2, 3

// ptr1 = head 1
// ptr2 = head.Next 1
