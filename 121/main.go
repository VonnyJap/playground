package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))

}

func maxProfit(prices []int) int {

	var minPrice = 1000000000
	var maxProfit = 0

	for _, p := range prices {
		if p < minPrice {
			minPrice = p
			continue
		}
		if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}

	return maxProfit
}

// given arr of prices[i] => stock prices day i to n
// purpose max profit => prices[i+j] - prices[i] => maximum

// one brute force solution that I can think of is
// for prices[i] => find j as such that prices[i+j] is maximum
// store this in the result arr for each [i]
// and then return the max

// we keep track of the min each time we found
// and each time the profit is back up - also calculate this
