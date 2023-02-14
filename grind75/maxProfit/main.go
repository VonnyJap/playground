package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}) == 5)
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}) == 0)

}

func maxProfit(prices []int) int {
	lowest := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < lowest {
			lowest = prices[i]
		}
		currentProfit := prices[i] - lowest
		if currentProfit > profit {
			profit = currentProfit
		}
	}

	return profit
}

// loop from 1 => n
// max profit = 0
// lowest = arr[0]
// if lowest found then we set that to the lowest
