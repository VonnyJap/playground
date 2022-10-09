package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3}, 10))
	fmt.Println(arraysIntersection([]int{197, 418, 523, 876, 1356}, []int{501, 880, 1593, 1710, 1870}, []int{521, 682, 1337, 1395, 1764}))
}

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {

	var result = []int{}
	for _, a := range arr1 {
		if binarySearch(arr2, a) && binarySearch(arr3, a) {
			result = append(result, a)
		}
	}

	return result
}

func binarySearch(a []int, x int) bool {

	start := 0
	end := len(a) - 1

	for start <= end {
		mid := (start + end) / 2
		if x == a[mid] {
			return true
		}
		if x > a[mid] {
			start = mid + 1
			continue
		}
		end = mid - 1
	}
	return false
}
