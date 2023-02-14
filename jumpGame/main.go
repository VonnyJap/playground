package main

import (
	"fmt"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
	fmt.Println(jump([]int{2, 3, 0, 1, 4}))

}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return jumpUtil(nums, 0, map[int]int{})
}

// nums fixed
// start is the changing input
func jumpUtil(nums []int, start int, memo map[int]int) int {
	if val, ok := memo[start]; ok {
		fmt.Println("memoized")
		return val
	}
	if start >= len(nums)-1 {
		memo[start] = 0
		return memo[start]
	}
	if nums[start] == 0 {
		memo[start] = 10000000000000
		return memo[start]
	}
	max := nums[start]
	result := []int{}
	for i := 1; i <= max; i++ {
		result = append(result, 1+jumpUtil(nums, start+i, memo))
	}

	memo[start] = min(result)

	return memo[start]
}

func min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	resp := arr[0]
	for _, a := range arr {
		if a < resp {
			resp = a
		}
	}
	return resp
}

// [2,3,1,1,4]
// 2
