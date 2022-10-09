package main

import "fmt"

func main() {

	lis := lisRecursive([]int{50, 3, 10, 7, 40, 80}, -1000000)
	fmt.Println(lis)
	fmt.Println(lisDP([]int{50, 3, 10, 7, 40, 80}))
}

// always start by making the problems smaller
// lis for [0,8,2] => 2
// going to try to use recursive first for this problem
// for the recursive - find the base case
// init a value of max that is possible for example starts with zero

func lisRecursive(seq []int, prev int) int {

	// find the base case first
	if len(seq) == 0 {
		return 0
	}

	count := 0
	if seq[0] > prev {
		count = 1 + lisRecursive(seq[1:], seq[0])
	}

	boom := lisRecursive(seq[1:], prev)

	if boom > count {
		return boom
	}
	return count
}

// lets do the dp for this

func lisDP(seq []int) int {

	var t []int

	for i := 0; i <= len(seq); i++ {
		t = append(t, 0)
	}
	t[0] = 1

	// start at the second element
	for i := 0; i < len(seq); i++ {
		for j := 0; j < i; j++ {
			if seq[j] < seq[i] && t[j] > t[i] {
				t[i] = t[j]
			}
		}
		t[i] = t[i] + 1
	}

	max := 0
	for _, v := range t {
		if v > max {
			max = v
		}
	}

	fmt.Println(t)
	return max
}
