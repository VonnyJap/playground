package main

import "fmt"

func main() {
	fmt.Println(countVowelStrings(1))
	fmt.Println(countVowelStrings(2))
	fmt.Println(countVowelStrings(33))
}

func countVowelStrings(n int) int {

	var result = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		result[i] = make([]int, 6)
	}

	for i := 0; i < 6; i++ {
		result[1][i] = i
	}

	for i := 2; i <= n; i++ {
		result[i][1] = 1
		for j := 2; j <= 5; j++ {
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}

	return result[n][5]
}
