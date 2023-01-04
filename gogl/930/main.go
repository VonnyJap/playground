package main

import "fmt"

func main() {
	fmt.Println(numSubarraysWithSum([]int{0, 1}, 1))

}

func numSubarraysWithSum(nums []int, goal int) int {

	// collect prefix sum
	// and then create a map out of it
	// loop through the prefix sum and if we can find prefix sum - goal => then the goal is valid
	// that means there is a prefix that can be minus off the current to get the goal

	return 0
}

// [1,0,1]
// pseudo code for recursive
// [1,0,1] -
// 1 - [0,1] 0
