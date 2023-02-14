package main

import "fmt"

func main() {
	fmt.Println(fib(10))
	fmt.Println(fib(15))

}

func fib(n int) int {
	memo := map[int]int{}
	return fibo(n, memo)
}

func fibo(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n <= 1 {
		memo[n] = n
		return memo[n]
	}
	minOne := fibo(n-1, memo)
	minTwo := fibo(n-2, memo)
	memo[n] = minOne + minTwo
	return memo[n]
}

// fibo is a 1d problem because only input is changing
// fibo with memo
// so how do we approach with memo
// store previous value in a hashmap
