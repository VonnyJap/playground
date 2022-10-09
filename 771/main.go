package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAbbbb"))
	fmt.Println(numJewelsInStones("z", "ZZ"))

}

func numJewelsInStones(jewels string, stones string) int {

	var dict = map[string]bool{}

	for _, j := range jewels {
		dict[string(j)] = true
	}

	var count = 0

	for _, s := range stones {
		if _, ok := dict[string(s)]; ok {
			count++
		}
	}
	return count
}

// use binary search to search and if found then increase ++
// also use hashmap to speed up
