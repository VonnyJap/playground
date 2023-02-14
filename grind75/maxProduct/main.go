package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-2, 3, -4}))
	fmt.Println(maxProduct([]int{2, -5, -2, -4, 3}))

}

// steps:
// 1. init a negativeValue = 1
// 2. max = nums[0]
// 3. loop
// - get negativeValue and store it here
// - still do the comparison accordingly
// - also with the product[i-1]*nums[i]

func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	minList := make([]int, len(nums))
	minList[0] = nums[0]
	maxList := make([]int, len(nums))
	maxList[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		arr := []int{minList[i-1] * nums[i], maxList[i-1] * nums[i], nums[i]}
		maxList[i] = max(arr)
		minList[i] = min(arr)
	}

	max := maxList[0]
	for i := 1; i < len(maxList); i++ {
		if max < maxList[i] {
			max = maxList[i]
		}
	}
	return max
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]

	for _, num := range arr[1:] {
		if num > max {
			max = num
		}
	}

	return max
}

func min(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	min := arr[0]

	for _, num := range arr[1:] {
		if num < min {
			min = num
		}
	}

	return min
}

// find how negative can it be in one loop
// after that build products
// so products are comparing 3 things
// 1. product[i-1]*nums[i], nums[i], negative[i-1]*nums[i]
// then get max from this
