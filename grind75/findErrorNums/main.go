package main

import "fmt"

func main() {
	fmt.Println(findErrorNums([]int{1, 2, 2, 4}))
}
func findErrorNums(nums []int) []int {

	cyclicSort(nums)
	result := []int{}

	for i, n := range nums {
		if nums[i] != i+1 {
			result = append(result, n)
			result = append(result, i+1)
			break
		}
	}
	return result
}

func cyclicSort(nums []int) {

	n := len(nums)
	pos := 0

	for i := range nums {
		nums[i]--
	}

	for pos < n {
		if nums[pos] < n && nums[pos] != nums[nums[pos]] {
			nums[pos], nums[nums[pos]] = nums[nums[pos]], nums[pos]
		} else {
			pos++
		}
	}

	for i := range nums {
		nums[i]++
	}
}
