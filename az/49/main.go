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

	var queue = strs[0:1]
	var done = map[string]bool{}
	var result = [][]string{}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if _, ok := done[current]; ok {
			continue
		}
		group := []string{current}
		done[current] = true
		origin := map[string]int{}
		for _, c := range current {
			if _, ok := origin[string(c)]; ok {
				origin[string(c)]++
				continue
			}
			origin[string(c)] = 1
		}
		for _, str := range strs {
			if _, ok := done[str]; ok {
				continue
			}
			target := map[string]int{}
			for _, c := range str {
				if _, ok := target[string(c)]; ok {
					target[string(c)]++
					continue
				}
				target[string(c)] = 1
			}
			if len(target) != len(origin) {
				queue = append(queue, str)
				continue
			}
			var good = true
			for k, v := range target {
				if _, ok := origin[k]; !ok {
					queue = append(queue, str)
					good = false
					break
				}
				if v != origin[k] {
					queue = append(queue, str)
					good = false
					break
				}
			}
			if good {
				group = append(group, str)
				done[str] = true
			}
		}
		result = append(result, group)
	}
	return result
}
