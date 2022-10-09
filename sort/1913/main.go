package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxProductDifference([]int{5, 6, 2, 7, 4}))
	fmt.Println(maxProductDifference([]int{4, 2, 5, 9, 7, 4, 8}))
}

func maxProductDifference(nums []int) int {

	// sort and then remove duplicate
	sort.Ints(nums)

	var unique = []int{nums[0]}

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			unique = append(unique, nums[i])
		}
	}

	if len(unique) <= 1 {
		return 0
	}

	return unique[len(unique)-1]*unique[len(unique)-2] - unique[0]*unique[1]
}
