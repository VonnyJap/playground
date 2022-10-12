package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArray([]int{1}))
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArrayDP([]int{5, 4, -1, 7, 8}))
	fmt.Println(maxSubArrayDP([]int{1}))
	fmt.Println(maxSubArrayDP([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArrayDP([]int{-2, -1}))

}

func maxSubArray(nums []int) int {

	var max int
	for i := 0; i < len(nums); i++ {
		var current int = 0
		for j := i; j < len(nums); j++ {
			current += nums[j]
			if max < current {
				max = current
			}
		}
	}
	return max
}

func maxSubArrayDP(nums []int) int {
	var result = make([]int, len(nums))
	result[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > result[i-1]+nums[i] {
			result[i] = nums[i]
			continue
		}
		result[i] = result[i-1] + nums[i]
	}

	var max = result[0]
	for _, m := range result {
		if max < m {
			max = m
		}
	}
	fmt.Println(nums)
	fmt.Println(result)

	return max
}

// optimize it
// 5, 4, -1
// if we implement something like result arr to hold the linear structure - call it L
// this is local maximum case
// L[0] = 0
// l[1] = which one is bigger L[0] + L[1] or L[1] => 5 => if nums[1] > nums[1] + L[0] => then L[1] = nums[1]
// else L[1] = nums[1] + L[0]
// l[2] = since L[1] is local maximum at L

// find largest of subarray
// subarray needs to be contiguous

// brute force
// 1. init a max value to hold the return result
// 2. for each i -> pair i with itself and all other elements after i -> for j = i up to j < len(nums)
// 3. and then we will compare with max on each iteration
