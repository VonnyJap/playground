package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2}))
}

func subsets(nums []int) [][]int {

	output := [][]int{}

	// for here we need to loop through from 0 up to len(nums)

	for i := 0; i <= len(nums); i++ {
		subsetsUtil(0, nums, []int{}, i, &output)
	}

	return output
}

func subsetsUtil(start int, nums []int, curr []int, k int, output *[][]int) {

	if len(curr) == k {
		*output = append(*output, curr)
		return
	}
	for i := start; i < len(nums); i++ {
		subsetsUtil(i+1, nums, append(curr, nums[i]), k, output)
	}
}
