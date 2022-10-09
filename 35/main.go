package main

import "fmt"

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

func searchInsert(nums []int, target int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {
		middle := (start + end) / 2
		if nums[middle] == target {
			return middle
		}
		if target < nums[middle] {
			end = middle - 1
			continue
		}
		start = middle + 1
	}

	return start
}
