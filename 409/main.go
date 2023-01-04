package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("a"))
	fmt.Println(longestPalindrome("abccccdd"))
}

// Palindrome: reverse is the same
// how can Palindrome get built?
// 1. need to have equal amounts of alphabets - so meaning anything that can be divided by 2
// 2. after collecting this - check how many leftover alphabets - if there is leftover than add 1
// 3. else no adding anything

func longestPalindrome(s string) int {

	dict := map[rune]int{}

	for _, c := range s {
		if _, ok := dict[c]; ok {
			dict[c]++
			continue
		}
		dict[c] = 1
	}

	total := 0

	for _, count := range dict {
		if count < 2 {
			continue
		}
		total += (count - count%2)
	}

	if total == len(s) {
		return total
	}
	return total + 1
}
