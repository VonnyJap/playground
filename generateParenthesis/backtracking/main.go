package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(3))

}

func generateParenthesis(n int) []string {
	return generateParenthesisUtil("", 0, 0, n)
}

func generateParenthesisUtil(prefix string, open, close, n int) []string {
	if len(prefix) == n*2 {
		return []string{prefix}
	}
	result := []string{}
	if open < n {
		result = append(result, generateParenthesisUtil(prefix+"(", open+1, close, n)...)
	}
	if close < open {
		result = append(result, generateParenthesisUtil(prefix+")", open, close+1, n)...)
	}
	return result
}

// backtracking steps:
// 1. init an open, close, prefix, n
// 2. if open < n, we can append more "("
// 3. if close < open, we can append more ")"
// 4. if len(prefix)==n*2, add that as result
