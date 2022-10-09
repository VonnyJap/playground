package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
	fmt.Println(sortSentence("Myself2 Me1 I4 and3"))
}

func sortSentence(s string) string {

	// split first
	// and then create a type
	var words = strings.Split(s, " ")
	sort.Slice(words, func(i, j int) bool {
		first := words[i]
		second := words[j]
		return first[len(first)-1] < second[len(second)-1]
	})

	for i, w := range words {
		words[i] = string(w[:len(w)-1])
	}

	return strings.Join(words, " ")
}
