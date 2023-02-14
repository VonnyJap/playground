package main

import "fmt"

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("["))

}

func isValid(s string) bool {

	stack := []rune{}

	for _, c := range s {

		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		concat := string([]rune{last, c})
		if concat == "()" || concat == "[]" || concat == "{}" {
			continue
		}
		return false
	}

	if len(stack) > 0 {
		return false
	}
	return true
}
