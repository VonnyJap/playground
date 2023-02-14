package main

import (
	"fmt"
)

func main() {
	fmt.Println(subsets([]int{0}))
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {

	result := [][]int{}

	for i := 0; i <= len(nums); i++ {
		result = append(result, subsetsUtil([]int{}, nums, i)...)
	}

	return result
}

func subsetsUtil(begin []int, nums []int, k int) [][]int {

	result := [][]int{}
	if len(begin) == k {
		result = append(result, begin)
		return result
	}

	for i := 0; i < len(nums); i++ {
		result = append(result, subsetsUtil(append(begin, nums[i]), nums[i+1:], k)...)
	}
	return result
}

// def subs(l):
//     if l == []:
//         return [[]]

//     x = subs(l[1:])

//     return x + [[l[0]] + y for y in x]
