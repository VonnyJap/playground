package main

import "sort"

func main() {}

func targetIndices(nums []int, target int) []int {
	// sort and then binary search

	sort.Ints(nums)

	var indices = []int{}
	for i, n := range nums {
		if n == target {
			indices = append(indices, i)
		}
	}

	return indices
}
