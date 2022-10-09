package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedInsert(head *ListNode, node *ListNode) *ListNode {
	if head == nil {
		return node
	}
	if head.Val >= node.Val {
		node.Next = head
		return node
	}
	var current = head
	for {
		if current.Next == nil || current.Next.Val > node.Val {
			break
		}
		current = current.Next
	}
	node.Next = current.Next
	current.Next = node
	return node
}
