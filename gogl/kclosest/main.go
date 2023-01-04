package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
	fmt.Println(findClosestElements([]int{2, 4, 5, 6, 9}, 3, 6))
}

func findClosestElements(arr []int, K, X int) []int {

	diff := make([]int, len(arr))

	for i, num := range arr {
		diff[i] = num - X
	}
	h := &IntHeap{}
	for _, d := range diff {
		heap.Push(h, d)
		if h.Len() > K {
			heap.Pop(h)
		}
	}

	result := []int{}
	for i := 0; i < K; i++ {
		diff := heap.Pop(h).(int)
		result = append(result, X+diff)
	}

	sort.Ints(result)
	return result
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return int(math.Abs(float64(h[i]))) > int(math.Abs(float64(h[j])))
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

// another way is to use the diff into heap
// we will take the min only
// we add the value to find it one by one

// iterate heap + target - we will get the value we want
