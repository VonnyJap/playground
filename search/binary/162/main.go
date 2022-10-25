package main

import "fmt"

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
	fmt.Println(findPeakElement([]int{1}))
	fmt.Println(findPeakElement([]int{1, 2}))

}

func findPeakElement(nums []int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {
		middle := (start + end) / 2
		if middle > 0 {
			if nums[middle-1] < nums[middle] {
				if middle >= len(nums)-1 {
					return middle
				}
				if nums[middle+1] < nums[middle] {
					return middle
				}
				start = middle + 1
				continue
			}
			end = middle - 1
			continue
		}
		start = middle + 1
	}

	return 0
}

// how can I find peak element?

// how to identify a peak:
// - middle element to be bigger than the number before
// 	- middle - 1 has to be smaller
// 	- if middle != len(nums)-1 then it has to be smaller as well
// how to update right/left?
// if found middle like that - just update the left?
// else updat the right?

// [1,2,3,1] start =0, end = 3, middle = 1 => left = middle + 1
// [1,2,3,1] left =2, end = 3, middle = 2 => return this

// [1,2,1,3,5,6,4] start=0, end=7, middle=3 => left = middle+1
// [1,2,1,3,5,6,4] start=4, end=7, middle=5 => return this

// [1,2] start = 0, end =1, middle = 0

// can use the pivot thing
// but the problem is left or right increment
