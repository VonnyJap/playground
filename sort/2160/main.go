package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumSum(4009))
}

func minimumSum(num int) int {

	var nums = []int{}

	for {
		nums = append(nums, num%10)
		num = int(num / 10)
		if num <= 0 {
			break
		}
	}

	sort.Ints(nums)

	return (nums[0]+nums[1])*10 + nums[2] + nums[3]
}
