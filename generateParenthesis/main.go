package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))
	// fmt.Println(generateParenthesis(3))

}

func generateParenthesis(n int) []string {
	output := []string{}
	generateParenthesisUtil2("(", n, &output)
	return output
}

func generateParenthesisUtil(prefix string, n int) []string {
	if len(prefix) == n*2 {
		if isValid(prefix) {
			return []string{prefix}
		}
		return []string{}
	}
	return append(generateParenthesisUtil(prefix+")", n), generateParenthesisUtil(prefix+"(", n)...)
}

func generateParenthesisUtil2(prefix string, n int, output *[]string) {
	if len(prefix) == n*2 {
		if isValid(prefix) {
			*output = append(*output, prefix)
			// fmt.Println(output)
			return
		}
		return
	}
	// fmt.Println("output1: ", output)
	generateParenthesisUtil2(prefix+")", n, output)
	// fmt.Println("output2: ", output)
	generateParenthesisUtil2(prefix+"(", n, output)
	// fmt.Println("output3: ", output)
	// return append(generateParenthesisUtil(prefix+")", n), generateParenthesisUtil(prefix+"(", n)...)
}

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}
	stack := []string{}

	for ctr := 0; ctr < len(s); ctr++ {
		if len(stack) == 0 {
			stack = append(stack, string(s[ctr]))
			continue
		}
		var parentheses = stack[len(stack)-1] + string(s[ctr])
		if parentheses == "()" || parentheses == "[]" || parentheses == "{}" {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, string(s[ctr]))
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
