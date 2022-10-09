package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 3, 5, 7}, 18))
}

func coinChange(coins []int, target int) int {
	if len(coins) == 0 {
		return 0
	}
	if target == 0 {
		return 0
	}
	min := 1000
	for _, v := range coins {
		if target-v >= 0 {
			count := 1 + coinChange(coins, target-v)
			if min > count {
				min = count
			}
		}

	}
	return min
}
