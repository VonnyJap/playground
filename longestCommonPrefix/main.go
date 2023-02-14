package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	shortest := strs[0]
	i := 0
	for i < len(shortest) {
		good := true
		for j := 1; j < len(strs); j++ {
			if shortest[i] != strs[j][i] {
				good = false
				break
			}
		}
		if !good {
			break
		}
		i++
	}

	if i-1 < 0 {
		return ""
	}
	return shortest[0:i]
}

// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// when the strs empty or len == 1 => return first strs as the longest

// else find the shortest str
// i can find a solution for n2
// with dp it is also n^2

// Let’s say you are given an integer array nums and an integer k. Return true if nums has a subarray where its length is at least two, and the sum of the elements in the subarray is a multiple of k
