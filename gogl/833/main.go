package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}))
	// fmt.Println(findReplaceString("abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}))
	fmt.Println(findReplaceString("vmokgggqzp", []int{3, 5, 1}, []string{"kg", "ggq", "mo"}, []string{"s", "so", "bfr"}))

	// fmt.Println(findReplaceString(
	// 	"ehvfwtrvcodllgjctguxeicjoudmxbevzrvravkidnricwsbnxmxvdckzahmqzbrlqugtmjvoqbxarmlgjeqcorhnodvnoqfomdp",
	// 	[]int{1, 31, 44, 70, 23, 73, 76, 92, 90, 86, 42, 4, 50, 17, 53, 20, 55, 15, 38, 64, 25, 9, 7, 68, 60, 88, 96, 47, 57, 34, 81, 78, 28},
	// 	[]string{"hvf", "vzr", "cw", "jvo", "jo", "qb", "ar", "noqf", "dv", "rh", "ri", "wt", "mx", "gux", "dc", "eic", "kz", "ct", "kidn", "lq", "ud", "odll", "vc", "tm", "qz", "no", "om", "bn", "ahm", "vra", "jeqco", "ml", "xb"},
	// 	[]string{"ajq", "zb", "r", "fai", "e", "zs", "io", "snxd", "nw", "oi", "ofb", "quq", "gj", "nsys", "dk", "sf", "muj", "ll", "hqx", "k", "n", "ptrya", "f", "qek", "u", "dhj", "e", "kr", "waj", "rvkr", "roaoeq", "mci", "djw"},
	// ))
	fmt.Println(findReplaceString("abcde", []int{2, 2}, []string{"cdef", "bc"}, []string{"f", "fe"}))

}

// eajqquqrfptryagjllnsyssfenmdjwezbrvkrvhqxofbrskrxgjvdkmujwajubrkugqekfaizsxiomcigroaoeqoidhjnwsnxdedp
// eajqquqrfptryagjllnsyssfenmdjwezbrvkrvhqxofbrskrxgjvdkmujwajubrkugqekfaizsxiecigroaoeqoidhjnwsnxdomdp

func findReplaceString(s string, indices []int, sources []string, targets []string) string {

	var intervals = [][]int{}

	for i, index := range indices {
		length := len(sources[i])
		if index+length > len(s) {
			continue
		}
		if s[index:index+length] == sources[i] {
			intervals = append(intervals, []int{index, index + length - 1, i})
		}
	}
	intervals = findNonOverlap(intervals)

	if len(intervals) == 0 {
		return s
	}
	replacements := []string{s[0:intervals[0][0]]}

	for i, interval := range intervals {
		replacements = append(replacements, targets[interval[2]])
		if i >= len(intervals)-1 {
			fmt.Println(interval)
			continue
		}
		start := intervals[i][1] + 1
		end := intervals[i+1][0]
		replacements = append(replacements, s[start:end])
	}
	replacements = append(replacements, s[intervals[len(intervals)-1][1]+1:len(s)])

	return strings.Join(replacements, "")
}

func findNonOverlap(arr [][]int) [][]int {

	if len(arr) == 0 {
		return [][]int{}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i][0] < arr[j][0] })
	start := arr[0][0]
	end := arr[0][1]

	result := [][]int{}

	for i := 1; i < len(arr); i++ {
		if arr[i][0] <= end {
			if arr[i][1] > end {
				end = arr[i][1]
			}
			continue
		}
		if start == arr[i-1][0] && end == arr[i-1][1] {
			result = append(result, arr[i-1])
		}
		start = arr[i][0]
		end = arr[i][1]
	}

	if start == arr[len(arr)-1][0] && end == arr[len(arr)-1][1] {
		result = append(result, arr[len(arr)-1])
	}

	return result
}

// how to detect overlap?

// brute force

// this is sorted - we need to know the index k so we know which target to go - we can insert a 3rd value there
// based on the indices and sources - will perform interval
// find non overlap interval - using the merge interval concept thing
// and then using the non overlap interval - we make the replacement
