package main

import "fmt"

func main() {
	fmt.Println(tribonacci(3))
}

func tribonacci(n int) int {

	var result = make([]int, n+1)

	result[0] = 0
	if n == 0 {
		return result[n]
	}

	result[1] = 1
	if n == 1 {
		return result[n]
	}

	result[2] = 1
	if n == 2 {
		return result[n]
	}

	for i := 3; i <= n; i++ {
		result[i] = result[i-3] + result[i-2] + result[i-1]
	}

	return result[n]
}
