package main

import "fmt"

func main() {

	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
	fmt.Println(reverseVowels("aA"))

}

// collect first all the vowels position
// and use two pointers to reverse
func reverseVowels(s string) string {

	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}

	runes := []rune{}

	for _, r := range s {
		runes = append(runes, r)
	}

	start := 0
	end := len(s) - 1

	for start < end {
		_, starVowel := vowels[runes[start]]
		_, endVowel := vowels[runes[end]]
		if !starVowel && endVowel {
			start++
			continue
		}
		if starVowel && !endVowel {
			end--
			continue
		}
		if starVowel && endVowel {
			runes[start], runes[end] = runes[end], runes[start]
		}
		start++
		end--
	}

	return string(runes)
}
