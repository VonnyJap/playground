package main

import "fmt"

func main() {
	fmt.Println(reverseString("abc"))
}

// this can also be done iteratively
func reverseString(s string) string {

	if len(s) <= 1 {
		return s
	}

	return reverseString(s[1:len(s)]) + string(s[0])

}

// there can be Divide-and-Conquer Solution

func reverseStringDC(s string) string {

	if len(s) <= 1 {
		return s
	}

	return reverseString(s[1:len(s)]) + string(s[0])

}
