package main

import "fmt"

func main() {
	// fmt.Println(int('1'))
	// fmt.Println(int('2') + int('6'))
	// fmt.Println(int('1') + int('6'))
	// fmt.Println(int('2') + int('7'))

	fmt.Println(numDecodings("126"))
	fmt.Println(numDecodings("06"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("12"))

}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	return numDecodingsUtil(s)
}
func numDecodingsUtil(s string) int {

	if len(s) == 0 {
		return 1
	}

	aValue := int('1')
	zValue := int('2') + int('6')

	result := 0
	runeValue := 0

	for i, c := range s {
		runeValue += int(c)
		if runeValue < aValue || runeValue > zValue {
			break
		}
		result += numDecodingsUtil(s[i+1:]) // 1 & 26 => 2 && 6 => 1
	}
	return result
}

// number has to be 1 or more and 26 or smaller
// this is going to be a recursive function
// in the function
// loop through the string each char
// if string is >=1 and <= 26 then we can continue the recursion
// else continue
// add to result
// breaking condition when string == ""
