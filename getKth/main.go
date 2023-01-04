package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(getKth(12, 15, 2))
	fmt.Println(getKth(7, 11, 4))

	// fmt.Println(getPower(12))
	// fmt.Println(getPower(13))

	// fmt.Println(getPower(14))
	// fmt.Println(getPower(15))

}

// we can use heap to do the sorting
// we can setup a util to get the power
func getKth(lo int, hi int, k int) int {

	h := &PowerHeap{}

	for i := lo; i <= hi; i++ {
		power := Power{
			Num:   i,
			Power: getPower(i),
		}
		heap.Push(h, power)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(Power).Num
}

type Power struct {
	Num   int
	Power int
}

type PowerHeap []Power

func (h PowerHeap) Len() int {
	return len(h)
}

// the one got popped out always the bigger one
func (h PowerHeap) Less(i, j int) bool {
	if h[i].Power == h[j].Power {
		return h[i].Num > h[j].Num
	}
	return h[i].Power > h[j].Power
}

func (h PowerHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PowerHeap) Push(x interface{}) {
	*h = append(*h, x.(Power))
}

func (h *PowerHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func getPower(x int) int {
	if x == 1 {
		return 0
	}
	if x%2 == 0 {
		return getPower(x/2) + 1
	}
	return getPower(3*x+1) + 1
}
