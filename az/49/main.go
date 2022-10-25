package main

import "fmt"

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{"a"}))
	fmt.Println(groupAnagrams([]string{"", ""}))

}

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// brute force way
// what are the companions of eat - tea, ate - how do we get this? - loop through the thing and check if good
// maintain a queue and visited map of all strs
// queue is init as strs[0:1]
// visited is init with strs[0] -> true
// and then the result that holds the group
// when queue is not 0
// loop through strs and check:
// - if visited do nothing
// - if not visited then:
// 	- check if it is anagram of current str
// 		- if yes - add to the result group
// 		- if not - add to the queue

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	tmp := map[[26]int][]string{}
	for _, s := range strs {
		chars := [26]int{}
		for _, c := range s {
			chars[c-'a']++
		}
		fmt.Println(chars)
		tmp[chars] = append(tmp[chars], s)
	}
	for _, v := range tmp {
		res = append(res, v)
	}
	return res
}

// we can sort the alphabets
// func groupAnagrams(strs []string) [][]string {
// 	sorted := map[string][]int{}
// 	res := [][]string{}
// 	for i, v := range strs {
// 			tmp := []byte(v)
// 			sort.Slice(tmp, func(a, b int) bool { return tmp[a] < tmp[b] })
// 			sorted[string(tmp)] = append(sorted[string(tmp)], i)
// 	}
// 	for _, v := range sorted {
// 			tmp := []string{}
// 			for _, index := range v {
// 					tmp = append(tmp, strs[index])
// 			}
// 			res = append(res, tmp)
// 	}
// 	return res
// }

// // handle case of same string?
// func groupAnagrams(strs []string) [][]string {

// 	// var queue = strs[0:1]
// 	var queue = []int{0}      // this is the index of the string and no the string itself
// 	var done = map[int]bool{} // change this to handle index instead of string
// 	var result = [][]string{}

// 	for len(queue) > 0 {
// 		current := queue[0]
// 		queue = queue[1:]
// 		if _, ok := done[current]; ok {
// 			continue
// 		}
// 		group := []string{current}
// 		done[current] = true
// 		origin := map[string]int{}
// 		for _, c := range current {
// 			if _, ok := origin[string(c)]; ok {
// 				origin[string(c)]++
// 				continue
// 			}
// 			origin[string(c)] = 1
// 		}
// 		for _, str := range strs {
// 			if _, ok := done[str]; ok {
// 				continue
// 			}
// 			target := map[string]int{}
// 			for _, c := range str {
// 				if _, ok := target[string(c)]; ok {
// 					target[string(c)]++
// 					continue
// 				}
// 				target[string(c)] = 1
// 			}
// 			if len(target) != len(origin) {
// 				queue = append(queue, str)
// 				continue
// 			}
// 			var good = true
// 			for k, v := range target {
// 				if _, ok := origin[k]; !ok {
// 					queue = append(queue, str)
// 					good = false
// 					break
// 				}
// 				if v != origin[k] {
// 					queue = append(queue, str)
// 					good = false
// 					break
// 				}
// 			}
// 			if good {
// 				group = append(group, str)
// 				done[str] = true
// 			}
// 		}
// 		result = append(result, group)
// 	}
// 	return result
// }
