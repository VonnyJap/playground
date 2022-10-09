package main

import "fmt"

func main() {

	fmt.Println(maxProfit([]int{2, 4, 1}))

}

// func maxProfit(prices []int) int {

// 	max := 0
// 	// at each day look at what price to buy and look at future price to sell
// 	for i, buy := range prices[0 : len(prices)-1] {
// 		// sort the prices and take the last one
// 		sells := mergesort(prices[i:])
// 		localMax := sells[len(sells)-1] - buy
// 		if localMax > max {
// 			max = localMax
// 		}
// 	}

// 	return max
// }

// func mergesort(arr []int) []int {
// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	l1 := mergesort(arr[:len(arr)/2])
// 	l2 := mergesort(arr[len(arr)/2:])
// 	return merge(l1, l2)
// }

// func merge(a []int, b []int) []int {
// 	var c []int

// 	i := 0
// 	j := 0
// 	for i < len(a) && j < len(b) {
// 		if a[i] < b[j] {
// 			c = append(c, a[i])
// 			i++
// 		} else {
// 			c = append(c, b[j])
// 			j++
// 		}
// 	}
// 	for ; i < len(a); i++ {
// 		c = append(c, a[i])
// 	}
// 	for ; j < len(b); j++ {
// 		c = append(c, b[j])
// 	}
// 	return c
// }

// what about using DP?

func maxProfit(prices []int) int {

	if len(prices) <= 1 {
		return 0
	}
	profit := 0
	min := prices[0]

	for _, v := range prices {
		localMax := v - min
		if localMax > profit {
			profit = localMax
		}
		if v < min {
			min = v
		}
	}
	return profit
}
