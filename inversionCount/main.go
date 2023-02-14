package main

import "fmt"

func main() {
	fmt.Println(inversionCount([]int{5, 4, 3}))
	fmt.Println(inversionCountBrute([]int{5, 4, 3}))

	fmt.Println(inversionCount([]int{5, 4, 3, 2, 1}))
	fmt.Println(inversionCountBrute([]int{5, 4, 3, 2, 1}))

	// fmt.Println(inversionCount([]int{1, 2, 3, 4, 5}))
	fmt.Println(inversionCount([]int{8, 4, 2, 1}))
	fmt.Println(inversionCountBrute([]int{8, 4, 2, 1}))

}

func inversionCountBrute(arr []int) int {
	inversionCount := 0
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				inversionCount++
			}
		}
	}
	return inversionCount
}

func inversionCount(arr []int) int {
	result, swap := mergeSort(arr)

	fmt.Println(result)
	return swap
}

func mergeSort(arr []int) ([]int, int) {

	if len(arr) < 2 {
		return arr, 0
	}
	total := 0
	l1, t1 := mergeSort(arr[:len(arr)/2])
	l2, t2 := mergeSort(arr[len(arr)/2:])

	result, t3 := merge(l1, l2)

	total = t1 + t2 + t3
	return result, total
}

func merge(a, b []int) ([]int, int) {
	var c []int

	swap := 0
	i := 0
	j := 0
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
			swap += len(a) - i
		}
	}
	for ; i < len(a); i++ {
		c = append(c, a[i])
	}
	for ; j < len(b); j++ {
		c = append(c, b[j])
	}
	fmt.Println("swap:", swap)

	return c, swap
}
