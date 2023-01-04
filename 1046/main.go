package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	fmt.Println(lastStoneWeight([]int{1}))
}

func lastStoneWeight(stones []int) int {
	h := &IntHeap{}

	for _, s := range stones {
		heap.Push(h, s)
	}

	for (*h).Len() > 1 {
		// pop two items
		y := heap.Pop(h).(int)
		x := heap.Pop(h).(int)
		if x == y {
			continue
		}
		heap.Push(h, y-x)
	}

	if len(*h) == 0 {
		return 0
	}

	return heap.Pop(h).(int)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
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
