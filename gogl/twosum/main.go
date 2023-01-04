package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 17))

}

func twoSum(arr []int, target int) []int {

	start := 0
	end := len(arr) - 1
	result := []int{}

	for start < end {
		if arr[start]+arr[end] == target {
			result = []int{start, end}
			break
		}
		if arr[start]+arr[end] > target {
			end--
			continue
		}
		start++
	}

	return result
}
