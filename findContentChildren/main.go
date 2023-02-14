package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}

// loop through using two pointers and see which one stopping first
func findContentChildren(g []int, s []int) int {

	i := 0
	j := 0

	happy := 0
	sort.Ints(g)
	sort.Ints(s)

	for i < len(g) && j < len(s) {

		if g[i] <= s[j] {
			s[j] = s[j] - g[i]
			i++
			happy++
		}
		j++
	}

	return happy
}
