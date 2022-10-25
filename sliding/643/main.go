package main

import "fmt"

func main() {

	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4))
	fmt.Println(findMaxAverage([]int{5}, 1))

}

func findMaxAverage(nums []int, k int) float64 {

	var maxSum int
	var maxAvg float64

	if len(nums) < k {
		return maxAvg
	}

	var tmpSum int
	var ctr = 0
	for ctr < k {
		tmpSum += nums[ctr]
		ctr++
	}

	maxSum = tmpSum

	for ctr < len(nums) {
		tmpSum = tmpSum + nums[ctr] - nums[ctr-k]
		if tmpSum > maxSum {
			maxSum = tmpSum
		}
		ctr++
	}

	maxAvg = float64(maxSum) / float64(k)

	return maxAvg
}
