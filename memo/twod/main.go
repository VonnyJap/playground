package main

import "fmt"

func main() {
	// fmt.Println(LCS("aa", "abd"))
	// fmt.Println(LCS("bcdaacd", "acdbac"))
	// fmt.Println(lcsRecur("aa", "abd"))
	// fmt.Println(lcsRecur("bcdaacd", "acdbac"))
	fmt.Println(lcsRecurWithMemo("aa", "abd", map[string]map[string]int{}))
	fmt.Println(lcsRecurWithMemo("bcdaacd", "acdbac", map[string]map[string]int{}))
	fmt.Println(lcsRecurWithMemo("AGGTAB", "GXTXAYB", map[string]map[string]int{}))

}

// what can we do with this?
// init an array of m+1 x n+1
// init all to zero
// and then we need to run two loops
// inside the loop -> compare the char
// - if same matrix[i-1][j-1] + 1
// - else max(matrix[i][j-1] + matrix[i-1][j])
// this is going to be bottom up approach

func longestCommonSubsequence(text1 string, text2 string) int {
	matrix := make([][]int, len(text1)+1)

	for i := range matrix {
		matrix[i] = make([]int, len(text2)+1)
	}

	for ai := range text1 {
		for bj := range text2 {
			if text1[ai] == text2[bj] {
				matrix[ai+1][bj+1] = 1 + matrix[ai][bj]
				continue
			}
			matrix[ai+1][bj+1] = max(matrix[ai][bj+1], matrix[ai+1][bj])
		}
	}

	return matrix[len(text1)][len(text2)]
}

func LCS(a, b string) int {

	matrix := make([][]int, len(a)+1)

	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
	}

	for ai := range a {
		for bj := range b {
			if a[ai] == b[bj] {
				matrix[ai+1][bj+1] = 1 + matrix[ai][bj]
				continue
			}
			matrix[ai+1][bj+1] = max(matrix[ai][bj+1], matrix[ai+1][bj])
		}
	}

	return matrix[len(a)][len(b)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// what about doing this recursive way

func lcsRecur(a, b string) int {

	if len(a) == 0 || len(b) == 0 {
		return 0
	}
	if a[0] == b[0] {
		return 1 + lcsRecur(a[1:], b[1:])
	}
	longerA := lcsRecur(a, b[1:])
	longerB := lcsRecur(a[1:], b)
	return max(longerA, longerB)
}

func lcsRecurWithMemo(a, b string, memo map[string]map[string]int) int {

	if _, ok1 := memo[a]; ok1 {
		if _, ok2 := memo[a][b]; ok2 {
			fmt.Println("memoized")
			return memo[a][b]
		}
	}
	if _, ok := memo[a]; !ok {
		memo[a] = map[string]int{}
	}
	// fmt.Println(memo)

	if len(a) == 0 || len(b) == 0 {
		memo[a][b] = 0
		// fmt.Println("bang:", memo)
		return memo[a][b]
	}
	if a[0] == b[0] {
		trailing := lcsRecurWithMemo(a[1:], b[1:], memo)
		memo[a][b] = 1 + trailing
		// fmt.Println("boom:", memo)
		return memo[a][b]
	}
	longerA := lcsRecurWithMemo(a, b[1:], memo)
	longerB := lcsRecurWithMemo(a[1:], b, memo)
	memo[a][b] = max(longerA, longerB)
	// fmt.Println("gong:", memo)
	return memo[a][b]
}
