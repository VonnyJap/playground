package main

import "fmt"

func main() {
	fmt.Println(validMountainArray([]int{2, 1}))
	fmt.Println(validMountainArray([]int{3, 5, 5}))
	fmt.Println(validMountainArray([]int{0, 3, 2, 1}))
}

func validMountainArray(arr []int) bool {

	// when foundPeak is not 1 then anything need to be strictly increasing
	if len(arr) < 3 {
		return false
	}

	start := 0
	for start < len(arr)-1 {
		// here we just make sure we have strictly increasing
		if arr[start] > arr[start+1] {
			break
		}
		start++
	}

	if start == 0 || start == len(arr)-1 {
		return false
	}

	for start < len(arr)-1 {
		// here we just make sure we have strictly increasing
		if arr[start] <= arr[start+1] {
			return false
		}
		start++
	}

	return true
}
