package main

import "fmt"

func main() {
	fmt.Println(fib(5))
}

func fib(n int) int {

	var fib = make([]int, n+1)

	// base cases
	fib[0] = 0
	if n == 0 {
		return fib[0]
	}

	fib[1] = 1
	if n == 1 {
		return fib[1]
	}

	for i := 2; i <= n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib[n]
}

// basically this is going to be recursive thing and i will do it using dp

// what is the base case?

// F(0) = 0, F(1) = 1

var temp = map[int]int{}

// func fib(n int) int {
// 	if _, ok := temp[n]; ok {
// 		return temp[n]
// 	}
// 	if n == 0 {
// 		return 0
// 	}
// 	if n == 1 {
// 		return 1
// 	}
// 	result := fib(n-1) + fib(n-2)
// 	temp[n] = result
// 	return result
// }
