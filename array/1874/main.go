package main

import "fmt"

func main() {
	fmt.Println(minProductSum([]int{2, 1, 4, 5, 7}, []int{3, 2, 4, 8, 6}))
}

func minProductSum(nums1 []int, nums2 []int) int {
	// sort both nums1 and nums2 and then multiply

	nums1 = mergesort(nums1)
	nums2 = mergesort(nums2) // reverse this

	fmt.Println(nums1)
	fmt.Println(nums2)

	sumprod := 0
	for i := range nums1 {
		sumprod = sumprod + nums1[i]*nums2[len(nums2)-1-i]
	}
	return sumprod
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
