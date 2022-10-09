package main

import "fmt"

func main() {
	fmt.Println(lcs("AB", "AB"))

}

func lcs(X, Y string) int {

	var result = make([][]int, len(X)+1)
	for i := 0; i < len(X)+1; i++ {
		result[i] = make([]int, len(Y)+1)
	}

	fmt.Printf("%+v\n", result)

	for i := 0; i < len(X); i++ {
		for j := 0; j < len(Y); j++ {
			if X[i] == Y[j] {
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

	fmt.Printf("%+v\n", result)

	return result[len(X)][len(Y)]
}
