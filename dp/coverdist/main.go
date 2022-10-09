package main

import "fmt"

func main() {
	fmt.Println(coverdist(2))
	fmt.Println(coverdist(3))
	fmt.Println(coverdist(4))
	fmt.Println(coverdist(7))
}

func coverdist(n int) int {

	if n < 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	// var result []int
	// for i := 1; i < n; i++ {
	// 	result = append(result, coverdist(n-i))
	// }

	// var final int

	// for _, r := range result {
	// 	final = final + r
	// }

	// return final

	return coverdist(n-1) + coverdist(n-2) + coverdist(n-3)
}

// for dp 3 base cases
// dp[0] = 1
// dp[1] = 1
// dp[2] = 2
