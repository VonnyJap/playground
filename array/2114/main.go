package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}))
}

func mostWordsFound(sentences []string) int {

	max := 0

	for _, s := range sentences {
		count := len(strings.Split(s, " "))
		if count > max {
			max = count
		}
	}

	return max
}
