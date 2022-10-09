package main

import "fmt"

func main() {
	X := "AGGTAB"
	Y := "GXTXAYB"
	fmt.Println(lcs(X, Y, len(X), len(Y)))
}

func lcs(arr1, arr2 string, m, n int) int {

	if m == 0 || n == 0 {
		return 0
	}
	if arr1[m-1] == arr2[n-1] {
		return 1 + lcs(arr1, arr2, m-1, n-1)
	}
	one := lcs(arr1, arr2, m, n-1)
	two := lcs(arr1, arr2, m-1, n)

	if one > two {
		return one
	}
	return two
}
