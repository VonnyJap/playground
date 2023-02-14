package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// max := 3
	// min := 0
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)
	// fmt.Println(rand.Intn(max-min) + min)

	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	Solution := Constructor(head)
	fmt.Println((&Solution).GetRandom())
	fmt.Println((&Solution).GetRandom())
	fmt.Println((&Solution).GetRandom())
	fmt.Println((&Solution).GetRandom())
	fmt.Println((&Solution).GetRandom())
	fmt.Println((&Solution).GetRandom())

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	elements []int
}

func Constructor(head *ListNode) Solution {

	sol := Solution{elements: []int{}}
	for head != nil {
		sol.elements = append(sol.elements, head.Val)
		head = head.Next
	}

	return sol
}

func (this *Solution) GetRandom() int {
	random := rand.Intn(len(this.elements))
	return this.elements[random]
}
