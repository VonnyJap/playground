package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstringKDistinct("ecebbaa", 2))
	fmt.Println(lengthOfLongestSubstringKDistinct("aa", 1))
	fmt.Println(lengthOfLongestSubstringKDistinct("eceba", 2))
	fmt.Println(lengthOfLongestSubstringKDistinct("s", 0))

	fmt.Println(lengthOfLongestSubstringKDistinct("bacc", 2))

}

// algo:
// each time len(seen) > k
// start increasing by 1
// and then map needs to be adjusted
// window is the same probably

func lengthOfLongestSubstringKDistinct(s string, k int) int {

	seen := map[string]int{}
	ctr := 0
	start := 0
	window := 0

	for ctr < len(s) {
		if _, ok := seen[string(s[ctr])]; !ok {
			seen[string(s[ctr])] = 1
		} else {
			seen[string(s[ctr])]++
		}
		// fmt.Println("one:", seen)
		if len(seen) > k {
			// fmt.Println("breaking")
			if seen[string(s[start])] == 1 {
				delete(seen, string(s[start]))
			} else {
				seen[string(s[start])]--
			}
			start++
		}
		// fmt.Println("two:", seen)
		ctr++
		if window < ctr-start {
			window = ctr - start
		}
	}

	return window
}

// init a map to contain the letter seen and running number
// run through the loop from i:=0 -> i<n
// what happening in the loop
// 1. check if the map has that letter/add if not
// 2. check if len of map > k
// break if that happens and reset the map
// if running number is smaller than current one -> update it? -> this is the window
