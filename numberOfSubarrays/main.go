package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
	fmt.Println(numberOfSubarrays([]int{2, 4, 6}, 1))
	fmt.Println(numberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2))

}

func numberOfSubarrays(nums []int, k int) int {

	prefix := make([]int, len(nums))
	odd := 0
	count := 0

	for _, n := range nums {
		prefix[odd]++
		if n%2 == 1 {
			odd++
		}
		if odd >= k {
			count += prefix[odd-k]
		}
	}

	return count
}

// using prefix sum
// odd = 0, prefix = make([]int, len(nums))

// nums = [1,1,2,1,1], k = 3
// 1 -> odd=0 -> prefix[0]++, odd=1
// [1,1] -> odd=1 -> prefix[1]++, odd=2
// [1,1,2] -> odd=2 -> prefix[2]++, odd=2
// [1,1,2,1] -> odd=2 -> prefix[2]++, odd=3
// count += prefix[odd-k]
// [1,1,2,1,1] -> odd=3 -> prefix[3]++, odd=4
// count += prefix[odd-k] prefix[4-3] -> prefix[1]
