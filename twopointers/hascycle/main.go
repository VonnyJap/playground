package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	// use two pointers for this case
	current := head
	next := head.Next

	for next != nil {
		next = next.Next
		if next == nil {
			return false
		}
		next = next.Next
		current = current.Next
		if next == current {
			return true
		}
	}

	return false
}
