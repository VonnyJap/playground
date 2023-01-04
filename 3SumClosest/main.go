package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	// fmt.Println(twoSumClosest([]int{1, 4, 6}, 11))
	// fmt.Println(twoSumClosest([]int{5, 4, 6}, 11))
	// fmt.Println(twoSumClosest([]int{0, 0, 0}, 1))
	// fmt.Println(twoSumClosest([]int{1, 2, 3, 4, 5}, 10))
	// fmt.Println(twoSumClosest([]int{-1, 2, 1, -4}, 4))

	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{0, 0, 0}, 1))

}

func threeSumClosest(nums []int, target int) int {

	sum := 0
	diff := 0

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			fmt.Println("here:", nums)
			twoSum := twoSumClosest(nums[j+1:], target-nums[j])
			if j == 0 {
				sum = twoSum + nums[j]
				diff = target - sum
				continue
			}
			newSum := twoSum + nums[j]
			newDiff := target - newSum
			if newDiff < diff {
				diff = newDiff
				sum = newSum
			}
		}
	}

	return sum
}

// we can do twoSum for each of the target probably?

// at the twoSum we can modify a little bit as such that

// keep that diff value and also an array to keep the result
// if we find something closer than before than we update

func twoSumClosest(nums []int, target int) int {

	fmt.Println("nums:", nums)
	fmt.Println("target:", target)

	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	copy := []int{}
	for _, n := range nums {
		copy = append(copy, n)
	}
	sort.Ints(copy)
	start := 0
	end := len(nums) - 1
	sum := 0

	for start < end {
		sum = nums[start] + nums[end]
		diff := int(math.Abs(float64(target) - float64(sum)))
		if diff == 0 {
			return sum
		}
		right := int(math.Abs(float64(target) - float64(nums[start]+nums[end-1])))
		if right < diff {
			end--
			continue
		}
		start++
	}

	return sum
}
