package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi(" -42"))
	fmt.Println(myAtoi("boom"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("20000000000000000000"))

}

func myAtoi(s string) int {

	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return 0
	}

	positive := true
	if trimmed[0] == '-' {
		positive = false
	}

	start := 0
	if trimmed[0] == '-' || trimmed[0] == '+' {
		start = 1
	}

	// and then read only the digit
	// once found non digit then break
	// add this to the queue and based on the length of digits keep on adding

	digits := []int{}

	for _, c := range trimmed[start:] {
		d := runeInt(c)
		if d >= 0 && d <= 9 {
			digits = append(digits, d)
			continue
		}
		break
	}

	atoi := 0

	for i, d := range digits {
		atoi += int(float64(d) * math.Pow(10, float64(len(digits)-1-i)))
	}

	if !positive {
		atoi *= -1
	}

	limit := int(math.Pow(2, 31))

	if positive {
		if atoi > limit-1 {
			atoi = limit - 1
		}
	} else {
		if atoi < limit*-1 {
			atoi = limit * -1
		}
	}

	return atoi
}

// [-231, 231 - 1]

func runeInt(r rune) int {
	return int(r - '0')
}
