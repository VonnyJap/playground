package main

import "fmt"

func main() {

	fmt.Println(backspaceCompare("ab#c", "ad#c"))
	fmt.Println(backspaceCompare("ab##", "c#d#"))
	fmt.Println(backspaceCompare("a#c", "b"))

	// fmt.Println(strWithoutBackspace("ab##"))
	// fmt.Println(strWithoutBackspace("c#d#"))
	// fmt.Println(strWithoutBackspace("a#c"))
	// fmt.Println(strWithoutBackspace("b"))

}

func backspaceCompare(s string, t string) bool {
	if strWithoutBackspace(s) == strWithoutBackspace(t) {
		return true
	}
	return false
}

func strWithoutBackspace(s string) string {

	ptr1 := 0
	ptr2 := 1

	for {
		// fmt.Println(s)
		// fmt.Println(ptr1)
		// fmt.Println(ptr2)

		// clean the first backspace
		if s == "" {
			break
		}
		if s[0] == '#' {
			s = s[1:]
			continue
		}
		if ptr1 < 0 || ptr2 > len(s)-1 {
			break
		}
		if s[ptr2] != '#' {
			ptr1++
			ptr2++
			continue
		}
		n := len(s)
		// else means ptr2 == '#'
		str := s[0:ptr1]
		if ptr2 < n-1 {
			str += s[ptr2+1:]
		}
		s = str
		if ptr1 == 0 {
			continue
		}
		ptr1--
		ptr2--
	}

	return s
}

// check one by one

// ab##

// ptr1 = 0, ptr2 = 1, no # found, ptr1++, ptr2++
// ptr1 = 1, ptr2 = 2, ptr2 = #, s = s[0:1]+s[3:] => a#, ptr1--, ptr2--
// ptr1 = 0, ptr2 = 1, ptr2 = #, s = s[0:0]+s[2:] // only if ptr2 not the last number
// breaking condition is when ptr1<0 or ptr2>len(s)-1

// what are the cases?
