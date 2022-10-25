package main

import "fmt"

func main() {
	fmt.Println(canConstruct("a", "b"))
	fmt.Println(canConstruct("aa", "ab"))
	fmt.Println(canConstruct("aa", "aab"))
}

// store magazine into the map
// and then loop ransomNote and check if the map is bigger than zero
func canConstruct(ransomNote string, magazine string) bool {

	var dict = map[rune]int{}

	for _, m := range magazine {
		if _, ok := dict[m]; ok {
			dict[m]++
			continue
		}
		dict[m] = 1
	}

	for _, r := range ransomNote {
		count, ok := dict[r]
		if !ok {
			return false
		}
		if count <= 0 {
			return false
		}
		dict[r]--
	}
	return true
}
