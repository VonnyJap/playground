package main

import "fmt"

func main() {
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(51))
	fmt.Println(convertToTitle(702))
	fmt.Println(convertToTitle(800))
}

// this is recursive problem
func convertToTitle(columnNumber int) string {

	if columnNumber <= 26 {
		return toCharStr(columnNumber)
	}

	var rem int = columnNumber % 26
	if rem == 0 {
		rem = 26
	}

	return convertToTitle(int(columnNumber-rem)/26) + toCharStr(rem)
}

func toCharStr(i int) string {
	return string(rune('A' - 1 + i))
}
