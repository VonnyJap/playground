package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProductDP([]int{2, 3, -2, 4}))
	fmt.Println(maxProductDP([]int{-2, 0, -1}))
	fmt.Println(maxProductDP([]int{-2, 3, -4}))
}

func maxProduct(nums []int) int {

	var product = nums[0]

	for i, _ := range nums {
		var current = 1
		for j := i; j < len(nums); j++ {
			current *= nums[j]
			if current > product {
				product = current
			}
		}
	}

	return product
}

// brute force
// 1. init a max value to hold the final result
// 2. for each element i : append element j where j = i up -> j < len(nums)
// 3. each time the interim product is more than max -> update max

// there are cases that can disrupt the chain for example 0 and negative numbers
// hence we need to keep track of all this
// 3 things that we need to compare:
// - current
// - max_so_far
// - min_so_far

func maxProductDP(nums []int) int {

	var maxLocal = nums[0]
	var minLocal = nums[0]
	var result = maxLocal

	for i := 1; i < len(nums); i++ {
		tempMax := max([]int{nums[i], nums[i] * maxLocal, nums[i] * minLocal})
		minLocal = min([]int{nums[i], nums[i] * maxLocal, nums[i] * minLocal})
		maxLocal = tempMax
		if maxLocal > result {
			result = maxLocal
		}
	}
	return result
}

func min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]
	for _, n := range arr {
		if n < min {
			min = n
		}
	}
	return min
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return max
}
