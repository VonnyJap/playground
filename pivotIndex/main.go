package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))

}

func pivotIndex(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	prefixes := make([]int, len(nums))
	prefixes[0] = nums[0]

	for i, num := range nums {
		if i == 0 {
			continue
		}
		prefixes[i] = prefixes[i-1] + num
	}

	for i := range nums {
		before := 0
		if i != 0 {
			before = prefixes[i-1]
		}
		if before == prefixes[len(prefixes)-1]-prefixes[i] {
			return i
		}
	}

	return -1
}
