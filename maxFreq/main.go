package main

import "fmt"

func main() {

	fmt.Println(maxFreq("aababcaab", 2, 3, 4))
	fmt.Println(maxFreq("aaaa", 1, 3, 3))
}

// this sounds like a sliding window thing

// lets implement brute force first
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	n := len(s)
	dict := map[string]int{}
	for i := 0; i <= n-minSize; i++ {
		for j := i + minSize; j <= n; j++ {
			start := s[i:j]
			if len(start) > maxSize {
				break
			}
			if _, ok := dict[start]; ok {
				dict[start]++
			} else {
				dict[start] = 1
			}
		}
	}
	max := 0
	for sub, count := range dict {
		if findNumLetters(sub) > maxLetters {
			continue
		}
		if count > max {
			max = count
		}
	}

	return max
}

func findNumLetters(s string) int {

	chars := map[rune]bool{}

	for _, c := range s {
		if _, ok := chars[c]; !ok {
			chars[c] = true
		}
	}

	return len(chars)
}
