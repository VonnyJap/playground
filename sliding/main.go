package main

func main() {}

func longestSubarray(nums []int, limit int) int {

	// bla bla
	// how to use sliding window here

	// so for each type of k then we move around?

	longest := 0
	// we update longest as we go around

	// when k == 1

	for k := 1; k <= len(nums); k++ {
		min, max := MinMax(nums[:k])
		delta := max - min
		if delta <= limit && k > longest {
			longest = k
		}
		for i := 0; i < len(nums)-k; i++ {

		}
	}

	return longest
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
