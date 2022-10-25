package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	fmt.Println(twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60))
	fmt.Println(twoSumLessThanK([]int{10, 20, 30}, 15))

}

func twoSumLessThanK(nums []int, k int) int {

	sort.Ints(nums)
	start := 0
	end := len(nums) - 1
	max := int(math.Inf(-1))

	for start < end {
		sum := nums[start] + nums[end]
		if sum < k {
			if max < sum {
				max = sum
			}
			start++
			continue
		}
		end--
	}

	if max != int(math.Inf(-1)) {
		return max
	}
	return -1
}

// what about doing brute force way?
// it will be O(n^2) - there will be two loops where i and then j starts from after i

// what about doing it using the two pointers way?
// nums = [34,23,1], k = 60
// steps:
// 1. sort [1,23,34]
// 2. start = 0, end = 2
// 3. while start < end then we will loop
// 4. in each loop - 1 + 34 -> 35 (less than 60) -> max = 35 -> start++
// 5. another loop - 23 + 34 -> 57 (less than 60) -> max = 57 -> start++
// 6. loop breaks
