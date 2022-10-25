package main

import "fmt"

func main() {
	fmt.Println(reorganizeString("aab"))
	fmt.Println(reorganizeString("aaab"))
}

// if the string length is odd - then the number of same chars at most will be n/2+1
// if the string length is even - then the number of same chars at most will be n/2
// we can loop through the string and store the chars count in a map

// "aab"
// - loop from i to len(s) -1
// 	- init a swapped bool = false
// 	- in each loop check if char is equal to char before
// 		- if yes then swap the current char with next char
// 			- if swapped is true - then break and return ""
// 		- else - swapped = false

func reorganizeString(s string) string {
	sByte := []byte(s)
	swapped := false
	for i := 1; i < len(sByte)-1; i++ {
		if sByte[i] != sByte[i-1] {
			swapped = false
			continue
		}
		// just got swapped
		if swapped {
			return ""
		}
		sByte[i], sByte[i+1] = sByte[i+1], sByte[i]
		swapped = true
	}
	if sByte[len(sByte)-1] == sByte[len(sByte)-2] {
		return ""
	}
	return string(sByte)
}
