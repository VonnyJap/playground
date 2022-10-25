package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))

}

func maxProduct(nums []int) int {

	var max int = -100000000000000000

	for i := 0; i < len(nums); i++ {
		localMax := nums[i]
		if localMax > max {
			max = localMax
		}
		for j := i + 1; j < len(nums); j++ {
			localMax = localMax * nums[j]
			if localMax > max {
				max = localMax
			}
		}
	}

	return max
}
