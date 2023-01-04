package main

import "fmt"

func main() {
	fmt.Println(maxSubarray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubarray([]int{5, 4, -1, 7, 8}))
}

// [-2,1,-3,4,-1,2,1,-5,4]
// -2 -> -2
// -1 vs 1 -> 1
// -3 vs -2 -> -2
// 4 vs 2 -> 4
// 3 vs -1 -> 3
// 2 vs 5 -> 5
// 6 vs 1 -> 6
// 1 vs -5 -> 1
// 5 vs 4 -> 5

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	sum := make([]int, len(nums))
	sum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if sum[i-1]+nums[i] > nums[i] {
			sum[i] = sum[i-1] + nums[i]
		} else {
			sum[i] = nums[i]
		}
	}

	max := sum[0]
	for i := 1; i < len(sum); i++ {
		if sum[i] > max {
			max = sum[i]
		}
	}

	return max
}
