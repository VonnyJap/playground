package main

import "fmt"

func main() {
	fmt.Println(pourWater([]int{1, 2, 3, 4}, 1, 2))
	fmt.Println(pourWater([]int{3, 1, 3}, 5, 1))
	fmt.Println(pourWater([]int{1, 2, 3, 4}, 2, 2))

	fmt.Println(pourWater([]int{2, 1, 1, 2, 1, 2, 2}, 4, 3))

}

// [3, 1, 3], 5, 1
// 1. pourWater([3, 2, 3], 4, 1)
// 2. pourWater([3, 3, 3], 3, 1)
// 3. pourWater([3, 4, 3], 2, 1)
// 3. pourWater([4, 4, 3], 1, 0)

// func pourWater(heights []int, volume int, k int) []int {

// 	if volume == 0 {
// 		return heights
// 	}

// 	// when it cannot move anywhere else then it should return back?

// 	// this should be something better
// 	if k == 0 || k == len(heights)-1 {
// 		return heights
// 	}

// 	// when we hit this volume is not 0

// 	// what about when k==0 or k==len(heights)-1

// 	currentVol := sum(heights)

// 	for {
// 		if sum(heights) >= currentVol+volume {
// 			break
// 		}
// 		if heights[k] > heights[k-1] {
// 			heights[k-1]++
// 			heights = pourWater(heights, volume-1, k-1)
// 		} else if heights[k] > heights[k+1] {
// 			heights[k+1]++
// 			heights = pourWater(heights, volume-1, k+1)
// 		} else {
// 			heights[k]++
// 			heights = pourWater(heights, volume-1, k)
// 		}
// 	}
// 	return heights
// }

// func sum(arr []int) int {

// 	sum := 0

// 	for _, a := range arr {
// 		sum += a
// 	}

// 	return sum
// }

func pourWater(heights []int, volume int, k int) []int {

	original := k

	for volume > 0 {
		// case when k == 0
		if k == 0 {
			if heights[k] > heights[k+1] {
				k++
				continue
			}
			heights[k]++
			volume--
			k = original
			continue
		}
		if k == len(heights)-1 {
			if heights[k] > heights[k-1] {
				k--
				continue
			}
			heights[k]++
			volume--
			k = original
			continue
		}

		// when we reach here we know that k cannot be 0 or len(heights)-1
		// when k is the lowest
		if heights[k] <= heights[k-1] && heights[k] <= heights[k+1] {
			heights[k]++
			volume--
			k = original
			continue
		}
		if heights[k] > heights[k-1] {
			k--
			continue
		}
		k++
	}
	return heights
}

// what will be a good iterative way to do this?

// def canMoveDirection(self, heights, currentLoc, K, direction):
// if currentLoc == -1 or currentLoc == len(heights):
// 		return [False, -1]
// if heights[currentLoc] < heights[K]:
// 		if currentLoc + direction >= 0 and currentLoc + direction < len(heights) and heights[currentLoc + direction] <= heights[currentLoc]:
// 				[canMove, loc] = self.canMoveDirection(heights, currentLoc + direction, K, direction)
// 				if canMove and heights[loc] < heights[currentLoc]:
// 						return [True, loc]
// 		return [True, currentLoc]
// elif heights[currentLoc] == heights[K]:
// 		return self.canMoveDirection(heights, currentLoc + direction, K, direction)
// else:
// 		return [False, -1]

// just do iter and not recur?

// 0. [1,2,3,4], v=2, k=2
// 1. [1,2,3,4], v=1, k=1
// 2. [2,2,3,4], v=0, k=1
// 3. [2,2,3,4], v=1, k=2
// 4. [2,3,3,4], v=0, k=2

// 0. [3,1,3], v=5, k=1 -> check if right or left of index is shorter -> if yes then the water starts dropping at the lower index - 1 droplet
// -> else the water stays there
// 	- update the heights accordingly as well
// - when water level next is lower
// 	- pourWater(heights, 1, k-1/k+1) => update heights with this value
// - what is the base case?
// 	-
// 1. [3,2,3], v=4, k=1
// 2. [3,3,3], v=3, k=1
// 3. [4,3,3], v=2, k=1
// 4. [4,4,3], v=1, k=1
// 5. [4,4,4], v=0, k=1
