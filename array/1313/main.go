package main

import "fmt"

func main() {
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4}))
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3}))
}

func decompressRLElist(nums []int) []int {

	var result = []int{}
	for i := 0; i < len(nums)/2; i++ {
		freq := nums[2*i]
		val := nums[2*i+1]
		for j := 0; j < freq; j++ {
			result = append(result, val)
		}
	}

	return result
}
