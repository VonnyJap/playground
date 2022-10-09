package main

import "fmt"

func main() {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {

	start := 0
	end := len(numbers) - 1

	var result = []int{}

	for start <= end {
		if numbers[start]+numbers[end] == target {
			result = []int{start + 1, end + 1}
			break
		}
		if numbers[start]+numbers[end] > target {
			end--
			continue
		}
		start++
	}

	return result

}
