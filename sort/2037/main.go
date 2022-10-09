package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMovesToSeat([]int{3, 1, 5}, []int{2, 7, 4}))
	fmt.Println(minMovesToSeat([]int{4, 1, 5, 9}, []int{1, 3, 2, 6}))
	fmt.Println(minMovesToSeat([]int{2, 2, 6, 6}, []int{1, 3, 2, 6}))
}

func minMovesToSeat(seats []int, students []int) int {

	var moves int

	// sort both and compare
	seats = mergesort(seats)
	students = mergesort(students)

	for i := range students {
		diff := students[i] - seats[i]
		if diff < 0 {
			diff *= -1
		}
		moves += diff
	}
	return moves

}

func minMovesToSeatUsingSort(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)
	var moves int

	for i := range students {
		diff := students[i] - seats[i]
		if diff < 0 {
			diff *= -1
		}
		moves += diff
	}
	return moves
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l1 := mergesort(arr[:len(arr)/2])
	l2 := mergesort(arr[len(arr)/2:])
	return merge(l1, l2)
}

func merge(a []int, b []int) []int {

	var result []int

	i := 0
	j := 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			result = append(result, a[i])
			i++
			continue
		}
		result = append(result, b[j])
		j++
	}

	for i < len(a) {
		result = append(result, a[i])
		i++
	}
	for j < len(b) {
		result = append(result, b[j])
		j++
	}

	return result
}
