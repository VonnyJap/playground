package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(beautySum("aabcb"))
	fmt.Println(beautySum("aabcbaa"))
}

func beautySum(s string) int {

	sum := 0

	for window := 3; window <= len(s); window++ {
		for pos := 0; pos <= len(s)-window; pos++ {
			diff := highLowDiff(s[pos : pos+window])
			sum += diff
		}
	}

	return sum

}

// steps:
// 1. group things by the map
// 2. loop the map and get min and max
// 3. return max - min => int
func highLowDiff(s string) int {

	// fmt.Println(s)
	dict := map[string]int{}
	for _, v := range s {
		if _, ok := dict[string(v)]; ok {
			dict[string(v)]++
			continue
		}
		dict[string(v)] = 1
	}
	// fmt.Println(dict)

	result := []int{}
	for _, v := range dict {
		result = append(result, v)
	}
	sort.Ints(result)
	// fmt.Println(result)

	return result[len(result)-1] - result[0]
}

// try len(s) == 4 => 0:3, 1:4

// work on the brute force solution
// get all possible substring - min should be three
// define a window for each substring and check the diff on each?
