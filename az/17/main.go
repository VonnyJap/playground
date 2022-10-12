package main

import "fmt"

func main() {
	one := letterCombinations("23")
	fmt.Println(one, len(one))
	two := letterCombinations("")
	fmt.Println(two, len(two))
	four := letterCombinations("235")
	fmt.Println(four, len(four))

}

var dict = map[string][]string{
	"1": alphabets(1, 26),
	"2": alphabets(1, 3),
	"3": alphabets(4, 6),
	"4": alphabets(7, 9),
	"5": alphabets(10, 12),
	"6": alphabets(13, 15),
	"7": alphabets(16, 19),
	"8": alphabets(20, 22),
	"9": alphabets(23, 26),
}

func alphabets(to, from int) []string {
	var result = []string{}
	for i := to; i <= from; i++ {
		result = append(result, fmt.Sprintf("%c", 'a'-1+i))
	}
	return result
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	if len(digits) == 1 {
		return dict[digits]
	}
	var result = []string{}
	later := letterCombinations(digits[1:])
	for _, c := range dict[digits[0:1]] {
		for _, l := range later {
			result = append(result, c+l)
		}
	}
	return result
}
