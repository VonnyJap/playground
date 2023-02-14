package main

import "fmt"

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
	fmt.Println(minNumberOfFrogs("crcoakroak"))
	fmt.Println(minNumberOfFrogs("croakcrook"))

	fmt.Println(minNumberOfFrogs("aoocrrackk"))

}

func minNumberOfFrogs(croakOfFrogs string) int {

	prefix := map[string]int{}
	word := "croak"

	for i := 0; i < len(word)-1; i++ {
		prefix[word[:i+1]] = 0
	}

	min := 0

	for _, char := range croakOfFrogs {

		var initial, concat string

		if char == 'c' {
			concat = "c"
		} else if char == 'r' {
			initial = "c"
			concat = "cr"
		} else if char == 'o' {
			initial = "cr"
			concat = "cro"
		} else if char == 'a' {
			initial = "cro"
			concat = "croa"
		} else if char == 'k' {
			initial = "croa"
		} else {
			return -1
		}

		if len(initial) > 0 {
			prefix[initial]--
		}
		if len(concat) > 0 {
			prefix[concat]++
		}

		total := 0
		for _, count := range prefix {
			if count < 0 {
				return -1
			}
			total += count
		}

		if min < total {
			min = total
		}
	}

	for _, count := range prefix {
		if count > 0 {
			return -1
		}
	}

	return min
}

// make use of hashmap and since "croak" has distinct chars
// we can make use of hashmap to see if such prefix exists

// loop through the word "croakcroak"
// each time that we cannot find the prefix - +1 for min

// map[string]int
// "c" -> map["c"] = 1 => min++
// "r" -> is there a map with key "c" (prefix)
// 	=> if yes then map["c"]=0, map["cr"]=1
// "o" -> map["cro"]=1, map[cr]=0
// "a" -> map["croa"]=1, map["cro"]=0
// "k" -> map["croak"]++, map["croa"]=0

// min = 0
// "crcoakroak"
// "c" -> map["c"]=1 => min=1
// "r" -> map["c"]=0; map["cr"]=1
// "c" -> map["c"]=1; map["cr"]=1
// "o" -> map["cr"]=0; map["cro"]=1
// "a" ->
