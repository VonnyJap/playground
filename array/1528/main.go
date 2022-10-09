package main

import "fmt"

func main() {

	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreString("abc", []int{0, 1, 2}))

}

func restoreString(s string, indices []int) string {

	result := make([]byte, len(s))
	for i, n := range indices {
		result[n] = s[i]
	}
	return string(result)
}
