package main

import "fmt"

func main() {
	fmt.Println(arrangeCoins(8))
}

func arrangeCoins(n int) int {

	var rows = 0

	for n > 0 {
		n = n - (rows + 1)
		if n >= 0 {
			rows++
		}
	}

	return rows
}
