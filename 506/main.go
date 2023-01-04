package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findRelativeRanks([]int{5, 4, 3, 2, 1}))
	fmt.Println(findRelativeRanks([]int{10, 3, 8, 9, 4}))

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

func findRelativeRanks(score []int) []string {

	// sort.Ints(score) - this will change the score and thats not what we want
	h := &IntHeap{}
	for _, val := range score { // O(N)
		heap.Push(h, val) // O(log K)
	}

	ctr := 1
	dict := map[int]string{}

	initialLen := h.Len()
	for i := initialLen; i > 0; i-- {
		val := heap.Pop(h).(int)

		if ctr == 1 {
			dict[val] = "Gold Medal"
		} else if ctr == 2 {
			dict[val] = "Silver Medal"
		} else if ctr == 3 {
			dict[val] = "Bronze Medal"
		} else {
			dict[val] = fmt.Sprintf("%d", ctr)
		}
		ctr++
	}

	// fmt.Println(dict)
	// fmt.Println(score)

	ranks := []string{}

	for _, s := range score {
		ranks = append(ranks, dict[s])
	}

	return ranks
}
