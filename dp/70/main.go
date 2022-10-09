package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {

	var result = make([]int, n+1)

	result[0] = 0
	if n == 0 {
		return result[n]
	}

	result[1] = 1
	if n == 1 {
		return result[n]
	}

	result[2] = 2
	if n == 2 {
		return result[n]
	}

	for i := 3; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}

	return result[n]
}
