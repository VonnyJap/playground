package main

import "fmt"

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
	fmt.Println(firstUniqChar("aabb"))
	fmt.Println(firstUniqChar("aadadaad"))

}

func firstUniqChar(s string) int {

	// use a map and store the value vs index
	// when finding same char then del the key
	dict := map[rune]int{}

	for _, c := range s {
		if _, ok := dict[c]; !ok {
			dict[c] = 1
			continue
		}
		dict[c]++
	}

	min := -1
	for i, c := range s {
		if dict[c] == 1 {
			return i
		}
	}

	return min
}
