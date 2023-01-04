package main

import (
	"container/heap"
	"fmt"
)

func main() {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(kWeakestRows(mat, 3))

	mat1 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	fmt.Println(kWeakestRows(mat1, 2))

}

// loop through mat and get the Mix

func kWeakestRows(mat [][]int, k int) []int {

	mix := &MixHeap{}

	for i, r := range mat {
		soldiers := 0
		for _, c := range r {
			if c == 1 {
				soldiers++
			}
		}
		heap.Push(mix, Mix{i, soldiers})
		if mix.Len() > k {
			heap.Pop(mix)
		}
	}

	return func() []int { // O (k log k)
		result := make([]int, mix.Len())
		initialLen := mix.Len()
		for i := initialLen; i > 0; i-- {
			m := heap.Pop(mix).(Mix)
			result[i-1] = m.Row
		}
		return result
	}()
}

type Mix struct {
	Row   int
	Count int
}

type MixHeap []Mix

func (h MixHeap) Len() int {
	return len(h)
}

func (h MixHeap) Less(i, j int) bool {
	if h[i].Count == h[j].Count {
		return h[i].Row > h[j].Row
	}
	return h[i].Count > h[j].Count
}

func (h MixHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MixHeap) Push(x interface{}) {
	*h = append(*h, x.(Mix))
}

func (h *MixHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
