package main

import "fmt"

func main() {
	fmt.Println(minAddToMakeValid("())"))
	fmt.Println(minAddToMakeValid("((("))
	fmt.Println(minAddToMakeValid("()))(("))
}

func minAddToMakeValid(s string) int {

	var stack = []string{}

	for _, c := range s {
		if len(stack) == 0 {
			stack = append(stack, string(c))
			continue
		}
		// append to the back and also grab from the back
		if stack[len(stack)-1]+string(c) == "()" {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, string(c))
	}

	return len(stack)
}
