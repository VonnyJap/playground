package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	i := 0
	j := 0

	result := []int{}

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
			continue
		}
		result = append(result, nums2[j])
		j++
	}

	for i < len(nums1) {
		result = append(result, nums1[i])
		i++
	}

	for j < len(nums2) {
		result = append(result, nums2[j])
		j++
	}

	if len(result) == 0 {
		return float64(0)
	}

	if len(result)%2 == 0 {
		n := len(result) / 2
		return (float64(result[n]) + float64(result[n-1])) / 2
	}

	return float64(result[len(result)/2])
}
