package main

import "fmt"

func main() {
	fmt.Println(getDescentPeriods([]int{3, 2, 1, 4}))
	fmt.Println(getDescentPeriods([]int{8, 6, 7, 7}))
	fmt.Println(getDescentPeriods([]int{1}))
}

func getDescentPeriods(prices []int) int64 {

	result := make([]int, len(prices))
	result[0] = 0

	for i := 1; i < len(prices); i++ {
		if prices[i-1]-prices[i] == 1 {
			result[i] = result[i-1] + 1
			continue
		}
		result[i] = 0
	}

	total := 0
	for _, r := range result {
		total += r
	}

	return int64(len(prices) + total)
}

// steps:
// 1. populate arr 0
// 2. at each loop
// 	- compare with the previous one and if differ by one then + 1 from previous
// 3. sum the result from loop and add the total numbers to get the count
