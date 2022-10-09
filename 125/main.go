package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))

}

func isPalindrome(s string) bool {

	sanitize := ""

	for _, c := range s {
		if checkAlphaChar(c) {
			sanitize += string(c)
			continue
		}
	}

	sanitize = strings.ToLower(sanitize)

	return sanitize == reverse([]byte(sanitize))
}

// a string is palindrome if after removing non alphanumberic when reverse read the same
// need to clean the string first
// and reverse it
// and compare the two

func checkAlphaChar(charVariable rune) bool {
	if (charVariable >= 'a' && charVariable <= 'z') || (charVariable >= 'A' && charVariable <= 'Z') || (charVariable >= '0' && charVariable <= '9') {
		return true
	}
	return false
}

func reverse(s []byte) string {

	start := 0
	end := len(s) - 1

	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}

	return string(s)
}
