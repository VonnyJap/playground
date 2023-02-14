package main

import "fmt"

func main() {

	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))

}

func missingNumber(nums []int) int {
	cyclicSort(nums)
	for i, _ := range nums {
		if nums[i] != i {
			return i
		}
	}
	return len(nums)
}

func cyclicSort(nums []int) {
	n := len(nums)
	pos := 0
	for pos < n {
		if nums[pos] < len(nums) && nums[pos] != nums[nums[pos]] {
			nums[pos], nums[nums[pos]] = nums[nums[pos]], nums[pos]
			continue
		}
		pos++
	}
}
