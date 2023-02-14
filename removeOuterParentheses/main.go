package main

import "fmt"

func main() {

	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}

func removeOuterParentheses(s string) string {

	primitive := ""
	primitiveList := []string{}

	for _, c := range s {
		primitive += string(c)
		if isValid(primitive) {
			primitiveList = append(primitiveList, primitive)
			primitive = ""
		}
	}

	result := ""
	for _, p := range primitiveList {
		result += p[1 : len(p)-1]
	}

	return result
}

// loop through s one by one

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
