package main

import "fmt"

func main() {
	fmt.Println(getConcatenation([]int{}))
}

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i, v := range nums {
		ans[i] = v
		ans[i+n] = v
	}
	return ans
}
