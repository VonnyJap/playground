package main

import "fmt"

func main() {
	fmt.Println(reverseArray([]int{10, 20, 30, 40, 50}))
}

func reverseArray(arr []int) []int {

	start, end := 0, len(arr)-1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

	return arr
}
