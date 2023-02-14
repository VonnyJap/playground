package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))

}

func productExceptSelf(nums []int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	right := make([]int, len(nums)+1)
	left := make([]int, len(nums)+1)

	left[0] = 1
	for i := 0; i < len(nums); i++ {
		left[i+1] = left[i] * nums[i]
	}

	right[len(right)-1] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = nums[i] * right[i+1]
	}

	// fmt.Println(left)
	// fmt.Println(right)

	result := []int{}

	for i := 0; i < len(nums); i++ {
		result = append(result, left[i]*right[i+1])
	}

	return result
}
