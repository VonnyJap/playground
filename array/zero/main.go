package main

import "fmt"

func main() {
	nums := []int{5, 0, 1, 2, 3, 4}
	fmt.Println(buildArray(nums))
}

func buildArray(nums []int) []int {
	var res []int
	for _, v := range nums {
		res = append(res, nums[v])
	}
	return res
}
