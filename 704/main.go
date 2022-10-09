package main

import "fmt"

func main() {

	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
}

func search(nums []int, target int) int {

	index := -1
	start := 0
	end := len(nums) - 1

	for start <= end {
		middle := (start + end) / 2
		if target == nums[middle] {
			index = middle
			break
		}
		if target < nums[middle] {
			end = middle - 1
			continue
		}
		start = middle + 1
	}

	return index
}
