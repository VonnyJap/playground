package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
	fmt.Println(reverseWords("God Ding"))
}

func reverseWords(s string) string {

	arr := strings.Split(s, " ")

	for i := range arr {
		arr[i] = string(reverseOne([]byte(arr[i])))
	}

	return strings.Join(arr, " ")
}

func reverseOne(s []byte) []byte {

	start := 0
	end := len(s) - 1

	for start <= end {
		s[start], s[end] = s[end], s[start]
		end--
		start++
	}

	return s
}
