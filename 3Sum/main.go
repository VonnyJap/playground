package main

import (
	"fmt"
	"sort"
)

func main() {

	// fmt.Println(threeSum([]int{-1, 0, 1}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}

func threeSum(nums []int) [][]int {

	final := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {

		sub := []int{}
		sub = append(sub, nums[i+1:]...)
		target := nums[i]
		result := twoSum(sub, target*(-1))
		if len(result) > 0 {
			final = append(final, []int{nums[i], sub[result[0]], sub[result[1]]})
		}
	}

	return final
}

func twoSum(numbers []int, target int) []int {

	start := 0
	end := len(numbers) - 1

	var result = []int{}

	for start < end {
		if numbers[start]+numbers[end] == target {
			result = []int{start, end}
			break
		}
		if numbers[start]+numbers[end] > target {
			end--
			continue
		}
		start++
	}
	return result
}

// to avoid duplicate - we can actually skip if prev == next

// brute force in n^2
// for each of nums make sure the rest are equal to that number?
