package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))

}

func findKthLargest(nums []int, k int) int {

	// use the int heap to work on this problem
	h := &IntHeap{}

	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	result := make([]int, h.Len())
	initialLen := h.Len()
	for i := initialLen; i > 0; i-- {
		result[i-1] = heap.Pop(h).(int)
	}
	return result[k-1]

}

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

// to solve this:
// 1. create a heap data structure: max heap?
// 2. each time
