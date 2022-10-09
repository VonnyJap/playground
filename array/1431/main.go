package main

import "fmt"

func main() {

	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {

	max := 0
	for _, c := range candies {
		if c > max {
			max = c
		}
	}

	result := []bool{}
	for _, c := range candies {
		if c+extraCandies >= max {
			result = append(result, true)
			continue
		}
		result = append(result, false)
	}
	return result
}

// get the highest number of candies and with that extraCandies add return the bool
