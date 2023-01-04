package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cababadd"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("eabcb"))
	fmt.Println(longestPalindromeCenter("babad"))
	fmt.Println(longestPalindromeCenter("cababadd"))
	fmt.Println(longestPalindromeCenter("cbbd"))
	fmt.Println(longestPalindromeCenter("eabcb"))
}

// use dp for working on this

func longestPalindrome(s string) string {

	n := len(s)
	table := make([][]int, n)

	for i := 0; i < n; i++ {
		table[i] = make([]int, n)
	}

	maxLength := 1

	for i := 0; i < n; i++ {
		table[i][i] = 1
	}

	// check for sub-string of length 2.
	start := 0
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			table[i][i+1] = 1
			start = i
			maxLength = 2
		}
	}

	// Check for lengths greater than 2.
	// k is length of substring
	for k := 3; k <= n; k++ {
		for i := 0; i < n-k+1; i++ {

			// Get the ending index of substring from
			// starting index i and length k
			j := i + k - 1

			// checking for sub-string from ith index to
			// jth index iff str.charAt(i+1) to
			// str.charAt(j-1) is a palindrome
			if table[i+1][j-1] == 1 && s[i] == s[j] {
				table[i][j] = 1
				if k > maxLength {
					start = i
					maxLength = k
				}
			}
		}
	}

	return s[start : start+maxLength]
}

// this sounds like the sliding window
func longestPalindromeCenter(s string) string {
	if len(s) <= 1 {
		return s
	}
	longest := s[0:1]

	for i := 0; i < len(s); i++ {
		temp := expand(s, i, i)
		if len(temp) > len(longest) {
			longest = temp
		}
		temp = expand(s, i, i+1)
		if len(temp) > len(longest) {
			longest = temp
		}
	}

	return longest
}

func expand(s string, begin, end int) string {
	for begin >= 0 && end <= len(s)-1 && s[begin] == s[end] {
		begin--
		end++
	}
	return s[begin+1 : end]
}
