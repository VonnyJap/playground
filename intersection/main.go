package main

import "fmt"

func main() {

	// head8 := &ListNode{Val: 1}
	// head8.Next = &ListNode{Val: 4}
	// head8.Next.Next = &ListNode{Val: 5}

	headA := &ListNode{Val: 3}
	// headA.Next = &ListNode{Val: 1}
	// headA.Next.Next = head8

	headB := &ListNode{Val: 4}
	// headB.Next = &ListNode{Val: 6}
	// headB.Next.Next = &ListNode{Val: 1}
	// headB.Next.Next.Next = head8

	fmt.Printf("%#v\n", getIntersectionNode(headA, headB))
	fmt.Printf("%#v\n", getIntersectionNodeFast(headA, headB))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// think about this iterative mode
func getIntersectionNode(headA, headB *ListNode) *ListNode {

	for headA != nil {
		b := headB
		for b != nil {
			if headA == b {
				return headA
			}
			b = b.Next
		}
		headA = headA.Next
	}

	return nil
}

// how do we solve with dp?

func getIntersectionNodeFast(headA, headB *ListNode) *ListNode {

	pa := headA
	pb := headB

	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
