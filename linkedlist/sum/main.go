package main

import "fmt"

func main() {

	list1 := &Node{Val: 4}
	list1.Next = &Node{Val: 5}
	list1.Next.Next = &Node{Val: 2}
	list2 := &Node{Val: 1}
	list2.Next = &Node{Val: 2}
	sum := add(list1, list2)
	fmt.Printf("%#v\n", sum.Val)
	fmt.Printf("%#v\n", sum.Next.Val)
	fmt.Printf("%#v\n", sum.Next.Next.Val)
	fmt.Printf("%#v\n", sum.Next.Next.Next)

}

type Node struct {
	Val  int
	Next *Node
}

func getSize(list *Node) int {
	size := 0
	for {
		if list == nil {
			return size
		}
		size++
		list = list.Next
	}
	return size
}

func add(list1, list2 *Node) *Node {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	size1 := getSize(list1)
	size2 := getSize(list2)

	if size1 == size2 {
		return addSameSize(list1, list2)
	}

	diff := size2 - size1
	if diff < 0 {
		list1, list2 = list2, list1
		diff = diff * (-1)
	}
	for i := 0; i < diff; i++ {
		list1 = &Node{
			Val:  0,
			Next: list1,
		}
	}
	// if not same size - make it same size by adding zero

	return addSameSize(list1, list2)
}

// in my implementation - i will put a carry here
// this can loop very deep inside so it will be able to add the least significant first

var carry = 0

func addSameSize(list1, list2 *Node) *Node {
	if list1 == nil {
		return list1
	}
	var new = &Node{}
	new.Next = addSameSize(list1.Next, list2.Next)
	sum := list1.Val + list2.Val + carry
	new.Val = sum % 10
	carry = int(sum / 10)

	return new
}
