package main

import "fmt"

func main() {
	s := []byte("hello")
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {

	start := 0
	end := len(s) - 1

	for start <= end {
		s[start], s[end] = s[end], s[start]
		end--
		start++
	}
}
