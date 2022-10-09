package main

import "fmt"

func main() {
	fmt.Println(lcs("XMJYAUZ", "MZJAWXU"))
}

func lcs(a, b string) int {

	var matrix [][]int
	for i := 0; i < len(a)+1; i++ {
		matrix = append(matrix, []int{})
		for j := 0; j < len(b)+1; j++ {
			matrix[i] = append(matrix[i], 0)
		}
	}
	for i := 1; i < len(a)+1; i++ {
		for j := 1; j < len(b)+1; j++ {
			if a[i-1] == b[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
				continue
			}
			if matrix[i-1][j] > matrix[i][j-1] {
				matrix[i][j] = matrix[i-1][j]
				continue
			}
			matrix[i][j] = matrix[i][j-1]
		}
	}
	return matrix[len(a)][len(b)]
}
