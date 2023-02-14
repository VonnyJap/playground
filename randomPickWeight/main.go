package main

import "math/rand"

func main() {}

type Solution struct {
	weights []int
	prefix  []int
	total   int
}

func Constructor(w []int) Solution {
	prefixList := []int{}
	prefix := 0
	for _, weight := range w {
		prefixList = append(prefixList, prefix+weight)
	}
	return Solution{
		weights: w,
		prefix:  prefixList,
		total:   prefix,
	}
}

func (this *Solution) PickIndex() int {
	random := rand.Intn(this.total-1) + 0
	for i, p := range this.prefix {
		if random < p {
			return i
		}
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
