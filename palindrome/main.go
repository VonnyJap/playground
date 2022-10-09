package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}

func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = strings.ToLower(reg.ReplaceAllString(s, ""))
	var new string

	for i, _ := range s {
		new = fmt.Sprintf("%s%s", new, string(s[len(s)-1-i]))
	}
	if new == s {
		return true
	}
	return false
}
