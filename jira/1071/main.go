package main

import "fmt"

func main() {
	fmt.Println(gcdOfStrings("LEETCODE", "LEET"))
}

// so we iterate until a == b
func gcdOfStrings(str1 string, str2 string) string {

	a := len(str1)
	b := len(str2)

	for a != b {
		if b > a {
			b = b - a
			continue
		}
		a = a - b
	}

	if str1[0:a] != str2[0:b] {
		return ""
	}

	prev := str1[0:a]

	for i := a; i < len(str1); i = i + a {
		if str1[i:a+i] != prev {
			return ""
		}
	}

	for i := a; i < len(str2); i = i + a {
		if str2[i:a+i] != prev {
			return ""
		}
	}

	return prev
}
