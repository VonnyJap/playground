package main

import "fmt"

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
	fmt.Println(peakIndexInMountainArray([]int{0, 5, 10, 2}))
}

func peakIndexInMountainArray(arr []int) int {

	if len(arr) < 3 {
		return 0
	}

	peakIndex := 0
	prev := arr[0]
	peak := prev
	window := 1

	for i := 1; i < len(arr); i++ {
		window++
		if arr[i] <= prev {
			if window >= 3 {
				if peak < prev {
					peak = prev
					peakIndex = i - 1
				}
				window = 1
			}
		}
		prev = arr[i]
	}

	return peakIndex

}

// [0,1,0]
// [0,2,1,0]
// [0,10,5,2]
