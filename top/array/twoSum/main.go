package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(numbers []int, target int) []int {

	// when sorting the numbers keep the index?

	numbersWithIndex := [][]int{}

	for i, n := range numbers {
		numbersWithIndex = append(numbersWithIndex, []int{i, n})
	}
	sort.Slice(numbersWithIndex, func(i, j int) bool {
		return numbersWithIndex[i][1] < numbersWithIndex[j][1]
	})

	// fmt.Println(numbersWithIndex)

	start := 0
	end := len(numbersWithIndex) - 1

	var result = []int{}

	for start < end {
		if numbersWithIndex[start][1]+numbersWithIndex[end][1] == target {
			result = []int{numbersWithIndex[start][0], numbersWithIndex[end][0]}
			break
		}
		if numbersWithIndex[start][1]+numbersWithIndex[end][1] > target {
			end--
			continue
		}
		start++
	}
	return result
}
