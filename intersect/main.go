package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	// use pointer in each nums1 and nums2
	i := 0
	j := 0

	result := []int{}

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
			continue
		}
		if nums1[i] < nums2[j] {
			i++
			continue
		}
		j++
	}

	return result
}
