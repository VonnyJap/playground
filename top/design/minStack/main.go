package main

import "fmt"

func main() {

	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))

}

func removeDuplicates(nums []int) int {

	removed := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		k := i - removed
		if nums[k] == nums[k+1] {
			nums = append(nums[:k+1], nums[k+2:]...)
			removed++
		}
	}

	return len(nums)
}
