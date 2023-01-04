package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))
	// fmt.Println(generateParenthesis(2))
	// fmt.Println(generateParenthesis(3))
	four := generateParenthesis(4)
	fmt.Println(four)
	fmt.Println(len(four))
}

// this can be done in recursion
// - when n==1 => "()"
// - preprocess
// 	1. call generateParenthesis based on n-1
// 	2. for each result returned
// 		- add "()" to the front -> check if exist in map
// 		- add "()" to the back -> check if exist in map
// 		- add "()" to the front and back -> check if exist in map
// 	3. then return

// then after that think about top down

// func generateParenthesis(n int) []string {

// 	if n == 1 {
// 		return []string{"()"}
// 	}

// 	pred := generateParenthesis(n - 1)

// 	dict := map[string]bool{}

// 	for _, p := range pred {
// 		if _, ok := dict[fmt.Sprintf("()%s", p)]; !ok {
// 			dict[fmt.Sprintf("()%s", p)] = true
// 		}
// 		if _, ok := dict[fmt.Sprintf("%s()", p)]; !ok {
// 			dict[fmt.Sprintf("%s()", p)] = true
// 		}
// 		if _, ok := dict[fmt.Sprintf("(%s)", p)]; !ok {
// 			dict[fmt.Sprintf("(%s)", p)] = true
// 		}
// 	}

// 	result := []string{}

// 	for k := range dict {
// 		result = append(result, k)
// 	}
// 	return result
// }

var result = []string{}

func generateParenthesis(n int) []string {
	result = []string{}
	dfs(0, 0, n, "")
	return result
}

func dfs(left, right, n int, s string) {

	if len(s) == n*2 {
		result = append(result, s)
		return
	}

	if left < n {
		dfs(left+1, right, n, s+"(")
	}
	if right < left {
		dfs(left, right+1, n, s+")")
	}
}
