package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3}))
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-1, -2, -3}))
	fmt.Println(maximumProduct([]int{-100, -98, -1, 2, 3, 4}))
}

func maximumProduct(nums []int) int {

	sort.Ints(nums)

	product := 1
	for i := 1; i <= 3; i++ {
		product *= nums[len(nums)-i]
	}

	negative := []int{}
	for _, n := range nums {
		if n >= 0 {
			break
		}
		negative = append(negative, n)
	}
	if len(negative) >= 2 {
		temp := negative[0] * negative[1] * nums[len(nums)-1]
		if temp > product {
			product = temp
		}
	}

	return product
}

// 1. sort the nums
// 2. check for negative and non negative
// 3. if non-negative >= 2 - take the two smallest
// 4. compare it with the current product so far
