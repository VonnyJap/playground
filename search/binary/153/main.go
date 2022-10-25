package main

import "fmt"

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}

func findMin(nums []int) int {
	pivot := findPivot(nums)
	// fmt.Println("pivot:", pivot)
	// fmt.Println("nums:", nums)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if pivot == 0 {
		if nums[pivot] < nums[pivot+1] {
			return nums[pivot]
		}
	}
	if pivot < len(nums) {
		pivot = pivot + 1
	}
	return nums[pivot]
}

func findPivot(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	start := 0
	end := len(nums) - 1
	if nums[start] < nums[end] {
		return 0
	}
	pivot := 0
	for start <= end {
		pivot = (start + end) / 2
		if pivot+1 < len(nums) {
			if nums[pivot] > nums[pivot+1] {
				break
			}
		}
		if nums[pivot] < nums[start] {
			end = pivot - 1
			continue
		}
		start = pivot + 1
	}
	return pivot
}
