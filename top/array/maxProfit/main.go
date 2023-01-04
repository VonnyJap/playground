package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))

}

func maxProfit(prices []int) int {

	profit := 0
	window := 1
	start := 0
	cum := 0

	for {
		if start+window >= len(prices) {
			cum = cum + profit
			break
		}
		newProfit := prices[start+window] - prices[start]
		if newProfit < profit {
			cum += profit
			start = start + window
			profit = 0
			window = 1
			continue
		}
		window++
		profit = newProfit
	}

	return cum
}

// 7,1,5,3,6,4
// 1. start=0, window=1, profit=0, cum=0, newProfit=1-7=-6 => cum=0, profit=0, start=1
// 2. start=1, window=1, profit=0, cum=0, newProfit=5-1=4 => cum=0, profit=4, window=2
// 3. start=1, window=2, profit=4, cum=0, newProfit=3-1=2 => cum=4, profit=0, start=1+2=3, window=1, profit=0
// 4. start=3, window=1, profit=0, cum=4, newProfit=6-3=3 => cum=4, profit=3, window=2
// 5. start=3, window=2, profit=3, cum=4, newProfit=4-3=1 => cum=4+3=7, profit=0, start=3+2, window=1
