package main

import "fmt"

func main() {
	// fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	// fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
	// fmt.Println(productExceptSelf([]int{0, -3, 3}))
	// fmt.Println(productExceptSelf([]int{0, 0}))
	// fmt.Println(productExceptSelf([]int{0, 4, 0}))
	fmt.Println(productExceptSelf([]int{4, 5, 1, 8, 2}))

}

// to approach this problem:
// 1. get the product of all elements in nums
// 2. to get the result list:
// 	- for each i: product/nums[i]

// func productExceptSelf(nums []int) []int {

// 	product := 0
// 	nonZero := []int{}
// 	for _, n := range nums {
// 		if n == 0 {
// 			continue
// 		}
// 		nonZero = append(nonZero, n)
// 	}

// 	if len(nonZero) > 0 {
// 		product = nonZero[0]
// 		for _, n := range nonZero[1:] {
// 			product *= n
// 		}
// 	}

// 	var result = make([]int, len(nums))

// 	for i := range result {
// 		if len(nonZero) == len(nums) {
// 			result[i] = product / nums[i]
// 			continue
// 		}
// 		if len(nums)-len(nonZero) == 1 {
// 			if nums[i] != 0 {
// 				result[i] = 0
// 			} else {
// 				result[i] = product
// 			}
// 		}
// 	}

// 	return result
// }

// using left and right product

func productExceptSelf(nums []int) []int {

	// product to the left of nums[i]
	left := make([]int, len(nums))
	// product to the right of nums[i]
	right := make([]int, len(nums))

	left[0] = 1
	right[len(right)-1] = 1

	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := 1; i < len(nums); i++ {
		right[len(right)-1-i] = right[len(right)-i] * nums[len(nums)-i]
	}
	fmt.Println(left)
	fmt.Println(right)

	var result = make([]int, len(nums))

	for i := range nums {
		result[i] = left[i] * right[i]
	}

	return result
}
