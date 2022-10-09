package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(2))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	mod := n % 2
	return mod == 0 && isPowerOfTwo(n/2)
}
