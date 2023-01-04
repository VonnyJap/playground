package main

import (
	"fmt"
	"math"
)

func main() {

	blocks := []map[string]bool{
		{
			"gym":    false,
			"school": true,
			"store":  false,
		},
		{
			"gym":    true,
			"school": false,
			"store":  false,
		},
		{
			"gym":    true,
			"school": true,
			"store":  false,
		},
		{
			"gym":    false,
			"school": true,
			"store":  false,
		},
		{
			"gym":    false,
			"school": true,
			"store":  true,
		},
	}

	reqs := []string{"gym", "school", "store"}

	fmt.Println(findMinDistance(blocks, reqs))

}

// at block 1
// 	1. gym(inf), school(0), store(inf) - inf
// 	2. gym(1), school(0), store(inf) - inf
// 	3. gym(1), school(0), store(inf) - inf
// 	4. gym(1), school(0), store(inf) - inf
// 	5. gym(1), school(0), store(4) - 4
// at block 2
// 	1. gym(inf), school(0), store(inf) - inf
// 	2. gym(1), school(0), store(inf) - inf
// 	3. gym(1), school(0), store(inf) - inf
// 	4. gym(1), school(0), store(inf) - inf
// 	5. gym(1), school(0), store(4) - 4

// what kind of structure then I need to tackle this?
// []map[string]int{} - map to do storing
// []int{} maxDistanceTraveled

func findMinDistance(blocks []map[string]bool, reqs []string) int {

	groups := []map[string]int{}

	for i, block := range blocks {
		groups = append(groups, map[string]int{})
		for place, exist := range block {
			if exist {
				groups[i][place] = 0
			}
		}
		if len(groups[i]) == len(reqs) {
			continue
		}
		for j := 0; j < len(blocks); j++ {
			if j == i {
				continue
			}
			currentBlock := blocks[j]
			for place, exist := range currentBlock {
				if exist {
					if _, ok := groups[i][place]; ok {
						if groups[i][place] > int(math.Abs(float64(j)-float64(i))) {
							groups[i][place] = int(math.Abs(float64(j) - float64(i)))
						}
						continue
					}
					groups[i][place] = int(math.Abs(float64(j) - float64(i)))
				}
			}
		}
	}

	maxDistanceTraveled := []int{}

	for _, group := range groups {
		dist := 0
		for _, placeDist := range group {
			if dist < placeDist {
				dist = placeDist
			}
		}
		maxDistanceTraveled = append(maxDistanceTraveled, dist)
	}

	minDist := maxDistanceTraveled[0]

	for i := 1; i < len(maxDistanceTraveled); i++ {
		if maxDistanceTraveled[i] < minDist {
			minDist = maxDistanceTraveled[i]
		}
	}

	return minDist
}
