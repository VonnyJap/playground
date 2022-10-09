package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))

}

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	var mapS = map[string]string{}
	var mapT = map[string]string{}

	i := 0
	for i < len(s) {
		ms, ok1 := mapS[string(s[i])]
		mt, ok2 := mapT[string(t[i])]

		if !ok1 && !ok2 {
			mapS[string(s[i])] = string(t[i])
			mapT[string(t[i])] = string(s[i])
			continue
		}

		if ms != string(t[i]) || mt != string(s[i]) {
			return false
		}
		i++
	}

	return true
}

// how many it differs and if those that got differs are in it
