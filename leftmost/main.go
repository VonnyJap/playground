package main

import "fmt"

func main() {
	fmt.Println(binarySearchFirstOne([]int{1, 1, 1, 1}))
}

type BinaryMatrix struct {
	Get        func(int, int) int
	Dimensions func() []int
}

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {

	dim := binaryMatrix.Dimensions()
	minCol := dim[1] // this will be equal to num cols to begin with - which should be last index + 1
	// for each col we will do binary search until we find the 1

	// for each of the row we will do binary search
	for i := 0; i < dim[0]; i++ {
		var values []int
		for j := 0; j < dim[1]; j++ {
			values = append(values, binaryMatrix.Get(i, j))
		}
		index := binarySearchFirstOne(values)
		if index >= 0 && index < minCol {
			minCol = index
		}
	}

	if minCol == dim[1] {
		return -1
	}
	return minCol
}

func binarySearchFirstOne(haystack []int) int {
	r := -1
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := (low + high) / 2
		if haystack[mid] == 1 && (mid == 0 || haystack[mid-1] == 0) {
			r = mid
			break
		}
		if haystack[mid] == 0 {
			low = mid + 1
			continue
		}
		high = mid - 1
	}

	return r
}
