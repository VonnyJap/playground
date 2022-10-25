package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumKeypresses("apple"))
	fmt.Println(minimumKeypresses("abcdefghijkl"))
}

func minimumKeypresses(s string) int {

	var dict = map[string]int{}
	for _, c := range s {
		if _, ok := dict[string(c)]; ok {
			dict[string(c)]++
			continue
		}
		dict[string(c)] = 1
	}

	arr := []int{}
	for _, n := range dict {
		arr = append(arr, n)
	}
	fmt.Println(arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)

	var press = 0

	for i, n := range arr {
		press += ((i / 9) + 1) * n
	}
	return press
}

// i know that i have 9 options
// i have to sort those and each time multiply by (divide by 9)
// i wanna sort reverse it
