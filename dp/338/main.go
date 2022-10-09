package main

import "fmt"

func main() {

	fmt.Println(countBits(1))
}

func countBits(n int) []int {

	var result = make([]int, n+1)

	result[0] = 0
	if n == 0 {
		return result
	}

	result[1] = 1
	if n == 1 {
		return result

	}

	for i := 2; i <= n; i++ {
		left := i % 2
		div := int(i / 2)
		result[i] = result[left] + result[div]
	}

	return result
}
