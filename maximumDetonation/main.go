package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}}))
	fmt.Println(maximumDetonation([][]int{{1, 1, 5}, {10, 10, 5}}))
	fmt.Println(maximumDetonation([][]int{{1, 2, 3}, {2, 3, 1}, {3, 4, 2}, {4, 5, 3}, {5, 6, 4}}))

}

// - endX=6, endY=6, currentStartX=

func maximumDetonation(bombs [][]int) int {

	sort.Slice(bombs, func(i, j int) bool {
		if bombs[i][0] == bombs[j][0] {
			return bombs[i][1] < bombs[j][1]
		}
		return bombs[i][0] < bombs[j][0]
	})

	return getOverlapCount(bombs)
}

// If C1C2 <= R1 – R2: Circle B is inside A.
// If C1C2 <= R2 – R1: Circle A is inside B.
// If C1C2 < R1 + R2: Circle intersects each other.
// If C1C2 == R1 + R2: Circle A and B are in touch with each other.
// Otherwise, Circle A and  do not overlap

func getOverlapCount(bombs [][]int) int {

	if len(bombs) <= 1 {
		return len(bombs)
	}

	max := 0
	overlap := 0

	for i := 1; i < len(bombs); i++ {

		current := bombs[i]
		prev := bombs[i-1]
		distance := int(math.Sqrt(math.Pow(float64((prev[0]-current[0])), 2) + math.Pow(float64((prev[1]-current[1])), 2)))
		diff := int(math.Abs(float64(current[2]) - float64(prev[2])))

		if distance <= diff || distance <= current[2]+prev[2] {
			overlap++
			continue
		}
		if overlap > max {
			max = overlap
		}
		overlap = 0
	}

	return overlap + 1
}
