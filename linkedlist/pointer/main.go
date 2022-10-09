package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func main() {

	temp := &Node{Value: 2}
	boom := &Node{Value: 10}
	n := temp
	n.Next = &Node{Value: 5}
	n.Value = 3
	fmt.Printf("%+v\n", *n)
	fmt.Printf("%+v\n", (*n).Next)

	fmt.Printf("%+v\n", *temp)

	*n = *boom
	fmt.Printf("%+v\n", *n)
	fmt.Printf("%+v\n", (*n).Next)
	fmt.Printf("%+v\n", *temp)

	var bang Node = *temp
	fmt.Printf("%+v\n", bang)
	bang.Value = 8
	fmt.Printf("%+v\n", bang)
	fmt.Printf("%+v\n", *temp)

}
