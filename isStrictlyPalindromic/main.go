package main

import "fmt"

func main() {
	fmt.Println(isStrictlyPalindromic(9))
	fmt.Println(isStrictlyPalindromic(4))

}

func isStrictlyPalindromic(n int) bool {

	for i := 2; i <= n-2; i++ {
		if !isBasePalindromic(n, i) {
			return false
		}
	}

	return true
}

func isBasePalindromic(n, base int) bool {

	var bases []int

	for n != 0 {
		bases = append(bases, n%base)
		n = n / 2
	}
	start := 0
	end := len(bases) - 1

	for start < end {
		if bases[start] != bases[end] {
			return false
		}
		start++
		end--
	}

	return true
}
