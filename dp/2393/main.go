package main

import "fmt"

func main() {

}

func countSubarrays(nums []int) int64 {
	n := len(nums)

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}

	fmt.Println(dp)

	sum := 0
	for _, d := range dp {
		sum += d
	}

	return int64(sum)
}

// def countSubarrays(self, nums):
// n = len(nums)

// dp = [1]*n

// for i in range(1,n):
// 		if nums[i] > nums[i-1]:
// 				dp[i] = dp[i-1] + 1

// return sum(dp)
