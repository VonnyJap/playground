package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 5, 5, 7, 8, 8, 8, 8, 10, 11, 11}, 5))
}

// so looking at this problem
// we can probably do a binary search modified?
// we need to find the search in one time
// and once we get that we check number before or after

// this func is the recursive pattern
// if this is recursive - pre, post, or in between?
// when do I need to stop? - when no more probably?
// think about this again tomorrow using recursive
func searchRange(nums []int, target int) []int {

	result := []int{-1, -1}
	index := binarySearch(nums, 0, len(nums)-1, target)
	// return if the target cannot be found
	if index == -1 {
		return result
	}

	index1 := index
	index2 := index
	result = []int{index1, index2}
	fmt.Println(result)

	for index1 > 0 && nums[index1-1] == target {
		fmt.Println("boom")
		index1 = binarySearch(nums, 0, index1-1, target)
		result[0] = index1
	}

	for index2 < len(nums)-1 && nums[index2+1] == target {
		fmt.Println("baam")
		index2 = binarySearch(nums, index2+1, len(nums)-1, target)
		result[1] = index2
	}
	// edge case
	// index => 0 and index => len(nums)-1

	// search for left first until cannot be found?

	// what if we do something like searching until the index is -1?

	return result
}

// this is a recursive binary search
func binarySearch(nums []int, start, end, target int) int {
	index := -1
	if len(nums) == 0 {
		return index
	}

	for start <= end {
		middle := (start + end) / 2
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

// [5,7,7,8,8,10] => 8
// what about do binary search and passing the array at the same time?
// 1. pass the array [-1, -1] into the binary search
// 2. case:
// - cannot find anything -> [-1,-1] is not changed
// - find something:
// - center -> recursive from the beginning to middle-1 and from middle + 1 to last
// - front -> changed the result[0] and only recursive from middle + 1 to last
// - last -> changed the result[1] and only recursive from beginning to middle-1
// 3. stop condition? - when there is nothing else can be found but that target
