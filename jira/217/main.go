package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	var dict = map[int]bool{}

	for _, num := range nums {
		if _, ok := dict[num]; ok {
			return true
		}
		dict[num] = true
	}
	return false
}
