package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))

}

// steps:
// 1. keep buying price
// 2. keep profit
// 3. return last element in the profit arr

func maxProfit(prices []int) int {

	buying := make([]int, len(prices))
	profit := make([]int, len(prices))

	buying[0] = prices[0]
	profit[0] = 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-buying[i-1] > profit[i-1] {
			profit[i] = prices[i] - buying[i-1]
		} else {
			profit[i] = profit[i-1]
		}
		if buying[i-1] > prices[i] {
			buying[i] = prices[i]
		} else {
			buying[i] = buying[i-1]
		}
	}

	// fmt.Println(buying)
	// fmt.Println(profit)

	return profit[len(profit)-1]
}
