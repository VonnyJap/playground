package main

import "fmt"

func main() {

}

type Node struct {
	value int
	next  *Node
}

type Stack struct {
	head *Node
	size int
}

// Size() function
func (s *Stack) Size() int {
	return s.size
}

// IsEmpty() function
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// Peek() function
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0
	}
	return s.head.value
}

// Push() function
func (s *Stack) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

// Pop() function
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("StackEmptyException")
		return 0, false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

// Print() function
func (s *Stack) Print() {
	temp := s.head
	fmt.Print("Values stored in stack are: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

// implement stack using slice in go
// stack is LIFO

// need to implement push and pop
