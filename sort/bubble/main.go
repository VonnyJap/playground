package main

import "fmt"

func main() {
	arr := []int{10, 14, 28, 11, 7, 16, 30, 50, 25, 18}
	fmt.Println(sort(arr))
}

func sort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {
		swapped := false
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				front := arr[j]
				back := arr[j+1]
				arr[j], arr[j+1] = back, front
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return arr
}
