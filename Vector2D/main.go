package main

import "fmt"

func main() {
	vec := Constructor([][]int{{1}, {}})
	fmt.Println((&vec).HasNext())
	fmt.Println((&vec).Next())
	fmt.Println((&vec).HasNext())
}

type Vector2D struct {
	elements [][]int
	row      int
	queue    []int
}

// edge case when the vec is empty
func Constructor(vec [][]int) Vector2D {
	if len(vec) == 0 {
		return Vector2D{vec, 0, []int{}}
	}
	return Vector2D{vec, 0, vec[0]}
}

// Next() - should give me next value
func (this *Vector2D) Next() int {
	for len(this.queue) == 0 {
		if this.row >= len(this.elements) {
			return 0
		}
		this.row++
		this.queue = this.elements[this.row]
	}
	next := this.queue[0]
	this.queue = this.queue[1:]
	return next
}

// HasNext() - should tell me something is next
// HasNext should not move the pointer
// run the same thing there but not advancing?
func (this *Vector2D) HasNext() bool {
	// check if this.queue is zero
	if len(this.queue) > 0 {
		return true
	}
	// else we need to increment the row but just for checking
	// when I hit here i know that I need to do row increment
	row := this.row + 1
	if row < len(this.elements) {
		if row != len(this.elements)-1 {
			return true
		}
		if len(this.elements[row]) == 0 {
			return false
		}
		return true
	}
	return false
}

// [[1, 2], [3], [4]]
// [null, 1, 2, 3, true, true, 4, false]

// ptr1: 0
// ptr2: 0

// what about using queue?

// this is for next
// [[], [1, 2], [3], [4]]
// ptr1: 0, ptr2: 0, queue: [1,2], next => queue[0], queue => queue[1:]
// 	- case 1: queue is actually empty
// 		- keep on advancing ptr1 until queue is not empty
// 	- case 2: queue is only 1
// 		- return queue[0] as next and ptr1++

// for HasNext
// 	1. check if the ptr1 is actually less than len(vec)
// 		- if yes -> true
// 	2. that means this is at the last row
// 		- check if queue is zero
// 			- if so return false
// 			- if no return true
