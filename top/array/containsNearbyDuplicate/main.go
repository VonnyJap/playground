package main

import "fmt"

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))

}

// use hashmap to do this
func containsNearbyDuplicate(nums []int, k int) bool {

	count := map[int][]int{}

	for i, n := range nums {
		if _, ok := count[n]; ok {
			if i-count[n][len(count[n])-1] <= k {
				return true
			}
			count[n] = append(count[n], i)
			continue
		}
		count[n] = []int{i}
	}

	return false
}
