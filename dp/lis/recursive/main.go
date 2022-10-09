package main

import "fmt"

func main() {

	arr := []int{1, 2, 7, 4}
	fmt.Println("result:", lis_global(arr, len(arr)))

}

func lis(arr []int, curr int) int {

	if curr == 0 {
		return 1
	}

	max := 1
	for i := curr - 1; i >= 0; i-- {
		if arr[i] < arr[curr] {
			m := lis(arr, i)
			if max < m+1 {
				max = m + 1
			}
		}
	}
	return max
}

func lis_global(arr []int, n int) int {

	if len(arr) == 0 {
		return 0
	}
	max := 1

	for i := range arr {
		m := lis(arr, i)
		if max < m {
			max = m
		}
	}
	return max
}
