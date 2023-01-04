package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
}

func numUniqueEmails(emails []string) int {
	// clean first
	// 1. remove (.)
	// 2. remove everything after (+)

	// and then use map to store the key value
	dict := map[string]bool{}

	for i, e := range emails {
		// fmt.Println("email:", e)
		atIndex := strings.Index(e, "@")
		// fmt.Println("atIndex:", atIndex)
		cleaned := strings.Replace(e[:atIndex], ".", "", -1)
		// fmt.Println("cleaned:", cleaned)
		plusIndex := strings.Index(cleaned, "+")
		// fmt.Println("plusIndex:", plusIndex)
		if plusIndex >= 0 {
			cleaned = cleaned[0:plusIndex]
			// fmt.Println("cleaned2:", cleaned)
		}
		emails[i] = cleaned + string(e[atIndex:])
		// fmt.Println("emails[i]:", emails[i])
		if _, ok := dict[emails[i]]; !ok {
			dict[emails[i]] = true
		}
	}

	// fmt.Println(dict)
	return len(dict)
}
