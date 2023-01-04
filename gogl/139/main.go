package main

import "fmt"

func main() {

	// fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	// fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	// fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	// fmt.Println(wordBreak("aaaaaaa", []string{"aaaa", "aaa"}))
	// fmt.Println(wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}))

	fmt.Println(wordBreakBFS("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreakBFS("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreakBFS("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
	fmt.Println(wordBreakBFS("aaaaaaa", []string{"aaaa", "aaa"}))
}

var memo = map[string]bool{}

func wordBreak(s string, wordDict []string) bool {

	dict := map[string]bool{}

	for _, word := range wordDict {
		dict[word] = true
	}

	return wordBreakRecur(s, 0, dict)
}

func wordBreakRecur(s string, start int, dict map[string]bool) bool {
	if start == len(s) {
		return true
	}
	if v, ok := memo[s[start:len(s)]]; ok {
		return v
	}
	for end := start + 1; end <= len(s); end++ {
		if _, ok := dict[s[start:end]]; ok {
			result := wordBreakRecur(s, end, dict)
			memo[s[end:len(s)]] = ok && result
			if result {
				return true
			}
		}
	}
	memo[s[start:len(s)]] = false
	return false
}

// leetcode

// 1. create a dictionary from wordDict
// 2. set ctr to 0 and len to 1, total -> 0 -> keep track the len found so far
// 3. loop
// 	- each loop: s -> s[ctr:ctr+len]
// 	- check if s can be found in dict
// 		- yes: len -> 1, position changed, total added - just use position
// 		- no: len++
// 	- ctr++
// 3. ctr++ and len++ until we exhaust the whole s
// 4. if total not equal to s then false, otherwise true

// word break needs to use recursion and memoization

// 1. need to find start and end for the check in dict
// 2. wb recurse start end until len(s)
// 3. need a place to store temp result
// 4. what happen in the wb recurse function
// 	- we have for loop to increment end until end < len(s)?
// - each loop check if the start + end exist in dict and the rest exist
// 	- if yes -> true
// 	- else -> false
// 5. base case -> when len(s) == 0 -> return true

// wordBreakBFS
// steps:
// 1. create dictionary
// 2. create a queue and the beginning of the queue will be the 0
// 3. run a loop when queue is not empty and each loop
// 	- check start to end if in map -> if so then add start pointer to queue
// 4. when end == length then return true
// 5. also I need to have visited map so not to repeat

func wordBreakBFS(s string, wordDict []string) bool {

	dict := map[string]bool{}

	for _, word := range wordDict {
		dict[word] = true
	}

	queue := []int{0}         // the queue is a collection of start index
	visited := map[int]bool{} // if a starting index is visited before - dont visit it again

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		if _, ok := visited[start]; ok {
			continue
		}
		for end := start + 1; end <= len(s); end++ {
			sub := s[start:end]
			if _, ok := dict[sub]; ok {
				queue = append(queue, end)
				if end == len(s) {
					return true
				}
			}
		}
		visited[start] = true
	}
	return false
}
