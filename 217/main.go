package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {

	var dict = map[int]int{}

	for _, n := range nums {
		if _, ok := dict[n]; ok {
			return true
		}
		dict[n]++
	}
	return false
}
