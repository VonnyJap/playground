package main

import "fmt"

func main() {
	fmt.Println(isAlienSorted([]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"))
}

func isAlienSorted(words []string, order string) bool {

	var dict = map[string]int{}

	// create dictionary or hash map kind so no need to loop
	for i, o := range order {
		dict[string(o)] = i
	}

	// start from words number 1
	// if previous less than the current - continue - if never break then means all good return true
	for i := 1; i < len(words); i++ {
		prev := words[i-1]
		current := words[i]
		// which one is longer?
		for len(prev) > 0 && len(current) > 0 {
			if dict[string(prev[0])] < dict[string(current[0])] {
				fmt.Println("boom")
				break
			}
			if dict[string(prev[0])] == dict[string(current[0])] {
				if len(prev) == 1 {
					prev = ""
				} else {
					prev = prev[1:]
				}
				if len(current) == 1 {
					current = ""
				} else {
					current = current[1:]
				}
				continue
			}
			return false
		}
		if len(prev) > 0 && len(current) == 0 {
			return false
		}
	}
	return true
}
