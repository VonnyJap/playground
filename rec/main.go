package main

import "fmt"

func main() {
	mutation := map[string]string{
		"b": "c",
		"y": "x",
		"a": "b",
		"c": "g",
	}

	fmt.Println(makeSequence(mutation, "a"))
	fmt.Println(makeSequence(mutation, "b"))
	fmt.Println(makeSequence(mutation, "y"))

}

func makeSequence(dict map[string]string, s string) []string {
	// base case
	char, ok := dict[s]
	if !ok {
		return []string{s}
	}
	arr := append([]string{s}, makeSequence(dict, char)...)
	return arr
}
