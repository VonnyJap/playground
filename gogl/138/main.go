package main

import "fmt"

func main() {

	head := &Node{Val: 3}
	head.Next = &Node{Val: 3}
	head.Next.Random = head
	head.Next.Next = &Node{Val: 3}

	copy := copyRandomList(head)
	head.Val = 5
	printList(copy)
}

func printList(n *Node) {
	for n != nil {
		fmt.Println("val:", n.Val)
		if n.Random != nil {
			fmt.Println("random:", n.Random.Val)
		}
		n = n.Next
	}
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {

	if head == nil {
		return head
	}

	nodes := []Node{}
	position := map[Node]int{}

	current := head
	ctr := 0
	for current != nil {
		nodes = append(nodes, *current)
		position[*current] = ctr
		current = current.Next
		ctr++
	}

	for i := 0; i < len(nodes); i++ {
		if nodes[i].Next != nil {
			nodes[i].Next = &(nodes[i+1])
		}
		if nodes[i].Random != nil {
			random := *(nodes[i].Random)
			pos := position[random]
			nodes[i].Random = &(nodes[pos])
		}
	}

	return &(nodes[0])
}

// how can I do a copy of this when I am only given a head

// Steps:
// 1. Traverse through the List and store into array by position
// 2. we can use hashmap to store val to position
// 3. Traverse through the array to get the deep copy
