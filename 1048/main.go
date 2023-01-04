package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(longestStrChain([]string{"a", "b", "ba", "bca", "bda", "bdca"}))
	// fmt.Println(longestStrChain([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}))
	// fmt.Println(longestStrChain([]string{"abcd", "dbqca"}))

	words := []string{"abcd", "abc", "bcd", "abd", "ab", "ad", "b"}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})
	fmt.Println(words)
}

func longestStrChain(words []string) int {

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return words[i] < words[j]
		}
		return len(words[i]) < len(words[j])
	})

	// then do lcs here
	result := make([]int, len(words))
	result[0] = 1

	for i := 1; i < len(words); i++ {

		// if current same with the before than result the same
		if len(words[i]) == len(words[i-1]) {
			result[i] = result[i-1]
			continue
		}
		common := lcs(words[i], words[i-1])
		if common == len(words[i-1]) {
			result[i] = result[i-1] + 1
			continue
		}
		result[i] = result[i-1]
	}

	return result[len(result)-1]
}

func lcs(X, Y string) int {

	var result = make([][]int, len(X)+1)
	for i := 0; i < len(X)+1; i++ {
		result[i] = make([]int, len(Y)+1)
	}

	// fmt.Printf("%+v\n", result)

	for i := 0; i < len(X); i++ {
		for j := 0; j < len(Y); j++ {
			if X[i] == Y[j] {
				result[i+1][j+1] = result[i][j] + 1
				continue
			}
			if result[i+1][j] > result[i][j+1] {
				result[i+1][j+1] = result[i+1][j]
				continue
			}
			result[i+1][j+1] = result[i][j+1]
		}
	}

	// fmt.Printf("%+v\n", result)

	return result[len(X)][len(Y)]
}

// // If last names are different, then use last
// // name to determine whether element i is less than
// // element j.
// if members[i].LastName != members[j].LastName {
// 	return members[i].LastName < members[j].LastName
// }
// // Otherwise, use first name to determine whether
// // element i is less than element j.
// return members[i].FirstName < members[j].FirstName
