package main

import "fmt"

func main() {
	fmt.Println(findTheWinner(6, 5))
}

func findTheWinner(n int, k int) int {

	var arr = []int{}

	for i := 1; i <= n; i++ {
		arr = append(arr, i)
	}

	// fmt.Println(arr)

	return findTheWinnerRec(n, k, arr)
}

func findTheWinnerRec(n int, k int, m []int) int {

	if n == 0 {
		return 0
	}
	// base case
	if n == 1 {
		return m[0]
	}

	var mod = k % n

	if mod == 0 { // that means the last one
		m = m[:n-1]
	} else {
		m = append(m[mod:], m[:mod-1]...)
	}

	// fmt.Println(m)
	return findTheWinnerRec(n-1, k, m)
}
