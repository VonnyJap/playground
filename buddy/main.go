package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("aa", "aa"))
	fmt.Println(buddyStringsBetter("aa", "aa"))

}

func buddyStrings(s string, goal string) bool {

	for i := range s {
		for j := i + 1; j < len(s); j++ {
			r := []rune(s)
			r[i], r[j] = r[j], r[i]
			if string(r) == goal {
				return true
			}
		}
	}

	return false

}

// this is almost right
// just need to think about the case
// for example when things are all same between s and goal
// s[i] == s[j] == goal[i] == goal[j]

func buddyStringsBetter(s string, goal string) bool {

	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		var seen = map[string]bool{}
		for _, char := range s {
			if _, ok := seen[string(char)]; ok {
				return true
			}
			seen[string(char)] = true
		}
		return false
	}

	var pairs [][]string

	for i := range s {
		if string(s[i]) != string(goal[i]) {
			pairs = append(pairs, []string{string(s[i]), string(goal[i])})
		}
	}

	if len(pairs) != 2 {
		return false
	}

	return pairs[0][1] == pairs[1][0] && pairs[0][0] == pairs[1][1]
}

// another approach is to do binary search for this
