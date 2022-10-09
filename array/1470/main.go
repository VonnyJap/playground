package main

import "fmt"

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}

func shuffle(nums []int, n int) []int {

	var result = []int{}

	for i := 0; i < n; i++ {
		result = append(result, []int{nums[i], nums[n+i]}...)
	}

	return result
}
