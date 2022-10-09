package main

import "fmt"

func main() {
	fmt.Println(maxDepth("(1+(2*3)+((8)/4))+1"))
	fmt.Println(maxDepth("(1)+((2))+(((3)))"))
	fmt.Println(maxDepth("(1)+((2))"))
}

func maxDepth(s string) int {

	// put everything into stack and each time the stack is empty we need to reset?

	var parentheses string
	for _, c := range s {
		if string(c) == "(" || string(c) == ")" {
			parentheses += string(c)
		}
	}
	if len(parentheses) == 0 {
		return 0
	}
	var max = 0
	var stack = []string{string(parentheses[0])}
	for i := 1; i < len(parentheses); i++ {
		if len(stack) == 0 {
			stack = append(stack, string(parentheses[i]))
			continue
		}
		if len(stack) > max {
			max = len(stack)
		}
		// when we reach here stack is not zero
		if stack[len(stack)-1]+string(parentheses[i]) == "()" {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, string(parentheses[i]))
	}

	return max
}
