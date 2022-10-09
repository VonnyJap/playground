package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {

	var stack []*ListNode

	for head != nil {
		stack = append(stack, head)
		head = head.Next
	}

	return stack[len(stack)/2]
}
