package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))

}

func rob(nums []int) int {
	return robWithMemo(nums, map[int]int{})
}

func robWithMemo(nums []int, record map[int]int) int {

	// what is the base case
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	if _, ok := record[len(nums)-1]; ok {
		return record[len(nums)-1]
	}
	r1 := robWithMemo(nums[0:len(nums)-2], record) + nums[len(nums)-1]
	r2 := robWithMemo(nums[0:len(nums)-1], record)

	if r1 > r2 {
		record[len(nums)-1] = r1
		return r1
	}
	record[len(nums)-1] = r2
	return r2
}
