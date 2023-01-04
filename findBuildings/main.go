package main

import "fmt"

func main() {
	fmt.Println(findBuildings([]int{4, 2, 3, 1}))
	fmt.Println(findBuildings([]int{4, 3, 2, 1}))
	fmt.Println(findBuildings([]int{1, 3, 2, 4}))

}

func findBuildings(heights []int) []int {

	total := len(heights)
	temp := make([]int, len(heights))
	temp[0] = heights[total-1]
	result := []int{total - 1}

	i := 1

	for i < total {
		index := total - 1 - i
		if heights[index] > temp[i-1] {
			temp[i] = heights[index]
			result = append([]int{index}, result...)
		} else {
			temp[i] = temp[i-1]
		}
		i++
	}

	return result
}

// brute force solution
// at each of the building - check the later buildings are shorter or no
// if yes add to the result
// if no skip
// and this is n^2 solution

// using dp bottom up approach
// 1. create the temp table to hold the value
// 2. flip the arr around so it starts from the back and the back always have an ocean view
// 3. in each iter
// - check if the current value is more than the prev one
// - if not then dont add and update the table to the previous value
