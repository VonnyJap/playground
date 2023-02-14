package main

import "fmt"

func main() {
	fmt.Println(pivotArray([]int{9, 12, 5, 10, 14, 3, 10}, 10))
	fmt.Println(pivotArray([]int{-3, 4, 3, 2}, 2))

}

func pivotArray(nums []int, pivot int) []int {

	before := []int{}
	after := []int{}

	// how many pivots - count by deduct
	for _, n := range nums {
		if n < pivot {
			before = append(before, n)
		}
	}

	for _, n := range nums {
		if n > pivot {
			after = append(after, n)
		}
	}

	deduct := len(nums) - (len(before) + len(after))

	for i := 1; i <= deduct; i++ {
		before = append(before, pivot)
	}

	return append(before, after...)
}
