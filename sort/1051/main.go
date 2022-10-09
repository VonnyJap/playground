package main

import (
	"fmt"
	"sort"
)

func main() {
	heightChecker([]int{1, 1, 4, 2, 1, 3})
}

func heightChecker(heights []int) int {

	expected := mergesort(heights)

	sort.Ints(expected)
	fmt.Println(expected)
	fmt.Println(heights)

	var diff = 0
	for i := range heights {
		if heights[i] != expected[i] {
			diff++
		}
	}

	return diff
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l1 := mergesort(arr[:len(arr)/2])
	l2 := mergesort(arr[len(arr)/2:])
	return merge(l1, l2)
}

func merge(a []int, b []int) []int {
	var c []int

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		c = append(c, a[i])
	}
	for ; j < len(b); j++ {
		c = append(c, b[j])
	}
	return c
}
