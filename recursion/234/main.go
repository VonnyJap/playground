package main

import "fmt"

func main() {
	head := &ListNode{Val: 1}
	// head.Next = &ListNode{Val: 2}
	// head.Next.Next = &ListNode{Val: 2}
	// head.Next.Next.Next = &ListNode{Val: 1}

	fmt.Println(isPalindrome(head))

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {

	if head == nil {
		return false
	}
	reverse := Reverse(head)
	for head != nil && reverse != nil {
		if head.Val != reverse.Val {
			return false
		}
		head = head.Next
		reverse = reverse.Next
	}
	if reverse != nil || head != nil {
		return false
	}
	return true
}

// use stack so it will not change the whole thing
func Reverse(list *ListNode) *ListNode {
	var stack []int
	var current *ListNode = list
	for current != nil {
		stack = append(stack, current.Val)
		current = current.Next
	}

	var result *ListNode = nil
	for _, s := range stack {
		prev := result
		result = &ListNode{Val: s}
		result.Next = prev
	}
	return result
}

func reverseRecursive(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var p = reverseRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

func printList(n *ListNode) {
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

// reverse the list and then loop through it
