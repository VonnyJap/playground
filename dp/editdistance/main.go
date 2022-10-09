package main

import "fmt"

func main() {
	fmt.Println(editDistance("food", "money", len("food"), len("money")))
}

func editDistance(str1, str2 string, m, n int) int {

	var result = make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		result[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if str1[i] == str2[j] {
				result[i+1][j+1] = result[i][j] + 1
				continue
			}
			if result[i+1][j] > result[i][j+1] {
				result[i+1][j+1] = result[i+1][j]
				continue
			}
			result[i+1][j+1] = result[i][j+1]
		}
	}

	diff := len(str2) - result[m][n]
	if diff < 0 {
		diff = diff * -1
	}
	return diff
}
