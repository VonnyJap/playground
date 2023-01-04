package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))

}

// sort the number
// in a loop check the diff and ensure it is 1 and not more

func missingNumber(nums []int) int {

	sort.Ints(nums)
	missing := 0
	n := len(nums)

	if nums[0] != 0 {
		return missing
	}
	if nums[n-1] != n {
		return n
	}

	for i := 1; i < n; i++ {
		if nums[i]-nums[i-1] == 2 {
			missing = nums[i-1] + 1
			break
		}
	}

	return missing
}
