package main

import "fmt"

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))

}

func smallerNumbersThanCurrent(nums []int) []int {
	// sort the number
	// create hashmap
	// and then loop through the nums
	sortedNums := mergesort(nums)

	var dict = map[int]int{}
	for i, n := range sortedNums {
		if _, ok := dict[n]; !ok {
			dict[n] = i
		}
	}

	var result = []int{}
	for _, n := range nums {
		result = append(result, dict[n])
	}

	return result
}

func mergesort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	l1 := mergesort(arr[:len(arr)/2])
	l2 := mergesort(arr[len(arr)/2:])
	return merge(l1, l2)
}

func merge(a, b []int) []int {
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
