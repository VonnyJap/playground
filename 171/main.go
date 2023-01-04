package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))

}

func runeNumber(r rune) int {
	return int(r) - 64
}

func titleToNumber(columnTitle string) int {

	num := 0

	for i := 0; i < len(columnTitle); i++ {
		num += runeNumber(rune(columnTitle[len(columnTitle)-1-i])) * int(math.Pow(float64(26), float64(i)))
	}

	return num

}

// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28

// 65-64 => 1 => A
