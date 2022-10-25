package main

import "fmt"

func main() {
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	// fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	// fmt.Println(search([]int{1}, 0))
	// fmt.Println(search([]int{1}, 1))
	// fmt.Println(search([]int{1, 3, 5}, 0))
	fmt.Println(search([]int{8, 9, 2, 3, 4}, 9))

}

func search(nums []int, target int) int {
	// basically we need to find the pivot first
	// once pivot is found then we can do the search
	pivot := findPivot(nums)
	fmt.Println(pivot)
	if target >= nums[0] && target <= nums[pivot] {
		fmt.Println("hey")
		return binarySearch(nums, 0, pivot, target)
	}
	fmt.Println("hoo")
	return binarySearch(nums, pivot+1, len(nums)-1, target)
}

// how to find the pivot
// it is about the same like binary search
// it started from the middle
// normally middle has to be bigger than left and smaller than right if sorted
// when do we find the pivot -> when pivot > pivot+1
// what happen when pivot > left, then left -> pivot + 1
// what happen when pivot < right, then right -> pivot - 1
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

// this is a recursive binary search
func binarySearch(nums []int, start, end, target int) int {
	index := -1
	if len(nums) == 0 {
		return index
	}
	// fmt.Println("start:", start)
	// fmt.Println("end:", end)

	for start <= end {
		middle := (start + end) / 2
		// fmt.Println("middle:", middle)
		if nums[middle] == target {
			index = middle
			break
		}
		if nums[middle] < target {
			start = middle + 1
			continue
		}
		end = middle - 1
	}

	return index
}
