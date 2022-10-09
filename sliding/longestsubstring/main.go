package main

import "fmt"

func main() {

	fmt.Println(lengthOfLongestSubstring("abcadd"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("pwwkepw"))

}

func lengthOfLongestSubstring(s string) int {

	if len(s) < 2 {
		return len(s)
	}

	var currSize = 1       // current size of the substring
	var maxSize = currSize // max size of the substring
	var start = 0          // start of the substring

	// hold the chars position when seen
	// first char is stored
	var dict = map[string]int{string(s[0]): 0}

	for i := 1; i < len(s); i++ {

		// fmt.Println("s i:", string(s[i]))
		pos, ok := dict[string(s[i])]
		// case not seen
		if !ok {
			// fmt.Println("case1")
			currSize++
		} else {
			// when seen and it is part of the substring
			if pos >= start {
				// fmt.Println("case2")
				if maxSize < currSize {
					maxSize = currSize
				}
				start = pos + 1
				currSize = i - start + 1
			} else {
				// fmt.Println("case3")
				currSize++
			}
		}
		dict[string(s[i])] = i
		// fmt.Println(dict)
	}

	if maxSize < currSize {
		maxSize = currSize
	}
	return maxSize
}
