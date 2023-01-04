package main

import (
	"container/heap"
	"fmt"
)

func main() {
	KthLargest := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(KthLargest.Add(3))
	fmt.Println(KthLargest.Add(5))
	fmt.Println(KthLargest.Add(10))
	fmt.Println(KthLargest.Add(9))
	fmt.Println(KthLargest.Add(4))

}

type KthLargest struct {
	h *IntHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {

	h := &IntHeap{}

	for _, n := range nums {
		heap.Push(h, n)
	}

	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)

	for len(*this.h) > this.k {
		heap.Pop(this.h)
	}

	x := heap.Pop(this.h).(int)
	heap.Push(this.h, x)

	return x
}

// need to pop and push again

// implement the heap first in golang

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
