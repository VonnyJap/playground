package main

import "fmt"

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}

// brute force way to do it is to keep track of a value that get changed when something deleted
func removeElement(nums []int, val int) int {

	deleted := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		j := i - deleted
		if nums[j] == val {
			deleted++
			nums = append(nums[:j], nums[j+1:]...)
		}
	}

	fmt.Println(nums)
	return len(nums)
}

// n = len(nums)
// i := 0; i < len(nums); i++
// [3,2,2,3], 3
// i=0, delete=0, j = i-delete = 0 -> i++, delete++, [2,2,3]
// i:=1, delete=1, j = i-delete = 0 -> i++
// i:=2, delete=1, j = i-delete = 1 -> i++
// i:=3, delete=1, j = i-delete = 2 -> i++, delete++

// return len(nums) after changing
