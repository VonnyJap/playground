package main

import "fmt"

func main() {
	fmt.Println(subsetXORSum([]int{1, 3}))
	fmt.Println(subsetXORSum([]int{5, 1, 6}))
	fmt.Println(subsetXORSum([]int{3, 4, 5, 6, 7, 8}))
}

func subsetXORSum(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	subset := findSubset(nums, 0, []int{})
	fmt.Println(subset)
	result := 0
	for _, sub := range subset {
		temp := 0
		for _, s := range sub {
			temp = temp ^ s
		}
		result += temp
	}

	return result
}

func findSubset(nums []int, index int, output []int) [][]int {

	result := [][]int{}
	if len(nums) == 0 {
		return result
	}
	if index == len(nums) {
		return append(result, output)
	}

	result = append(result, findSubset(nums, index+1, append(output, []int{}...))...)
	result = append(result, findSubset(nums, index+1, append(output, []int{nums[index]}...))...)

	return result

}
