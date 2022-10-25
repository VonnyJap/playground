package main

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
	fmt.Println(findAnagrams("abab", "ab"))
}

func findAnagrams(s string, p string) []int {

	result := []int{}
	count := len(p)
	pdict := map[rune]int{}
	for _, pr := range p {
		if _, ok := pdict[pr]; ok {
			pdict[pr]++
			continue
		}
		pdict[pr] = 1
	}

	sdict := map[rune]int{}

	for i := 0; i <= len(s)-count; i++ {
		if len(sdict) == 0 {
			for _, sr := range s[0:count] {
				if _, ok := sdict[sr]; ok {
					sdict[sr]++
					continue
				}
				sdict[sr] = 1
			}
		} else {
			sdict[rune(s[i-1])]--
			if sdict[rune(s[i-1])] == 0 {
				delete(sdict, rune(s[i-1]))
			}
			if _, ok := sdict[rune(s[i+count-1])]; ok {
				sdict[rune(s[i+count-1])]++
			} else {
				sdict[rune(s[i+count-1])] = 1
			}
		}

		if checkSame(sdict, pdict) {
			result = append(result, i)
		}
	}

	return result
}

func checkSame(sdict, pdict map[rune]int) bool {

	if len(sdict) != len(pdict) {
		return false
	}

	for key, value := range pdict {
		if sv, ok := sdict[key]; ok {
			if sv != value {
				return false
			}
			continue
		}
		return false
	}

	return true
}

// Input: s = "cbaebabacd", p = "abc"
// Output: [0,6]
// Explanation:
// The substring with start index = 0 is "cba", which is an anagram of "abc".
// The substring with start index = 6 is "bac", which is an anagram of "abc".

// step:
// 1. len of p for the sliding
// 2. map of p
// 3. slide each time and check if anagram by checking if map contains all zero
