package main

import "fmt"

func main() {
	fmt.Println(sortedSquares([]int{-3, 1, 2}))

}

func sortedSquares(nums []int) []int {

	start := 0
	end := len(nums) - 1

	var result = []int{}

	for start <= end {

		num1 := nums[start]
		if num1 < 0 {
			num1 *= -1
		}
		num2 := nums[end]
		if num2 < 0 {
			num2 *= -1
		}
		if num1 > num2 {
			result = append([]int{num1 * num1}, result...)
			start++
			continue
		}
		result = append([]int{num2 * num2}, result...)
		end--
	}
	return result
}
