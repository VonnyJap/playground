package main

import "fmt"

func main() {
	fmt.Println(divisorGame(2))
	fmt.Println(divisorGame(3))
	fmt.Println(divisorGame(4))

	// fmt.Println(divisorGame(5))
}

func divisorGame(n int) bool {
	return divisorGameUtil(n, "alice")
}

func divisorGameUtil(n int, player string) bool {

	found := false
	for i := 1; i < n; i++ {
		if n%i == 0 {
			n = n - i
			found = true
			break
		}
	}

	if !found {
		return player == "bob"
	}

	if player == "alice" {
		return divisorGameUtil(n, "bob")
	}
	return divisorGameUtil(n, "alice")
}

// this is a recursion problem
