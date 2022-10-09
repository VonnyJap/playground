package main

func main() {

}

// use stack
// I will loop and build a stack
// two cases here to solve
// when we find a pair like "()", "[]", "{}"
// then we will remove one item from stack
// else
// we will add to the stack

// if len(stack) == 0 then true
// else false

// Input: s = "()"
// Output: true

// Input: s = "()[]{}"
// Output: true

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
