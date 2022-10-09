package main

import "fmt"

func main() {
	head := ImmutableListNode{}
	// head.Next = &ImmutableListNode{Val: 2}
	// head.Next.Next = &ImmutableListNode{Val: 3}
	// head.Next.Next.Next = &ImmutableListNode{Val: 4}
	// head.Next.Next.Next = &ImmutableListNode{Val: 4}

	printLinkedListInReverse(head)
}

type ImmutableListNode struct {
	Val  int
	Next *ImmutableListNode
}

func (this *ImmutableListNode) getNext() ImmutableListNode {
	return *(this.Next)
}

func (this *ImmutableListNode) printValue() {
	fmt.Println(this.Val)
}

func printLinkedListInReverse(head ImmutableListNode) {

	var current = head
	var next = current.getNext()
	if &next == nil {
		current.printValue()
		return
	}
	printLinkedListInReverse(next)
	current.printValue()
}
