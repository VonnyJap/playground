// package main

// This was for an mid-level role at Roblox.

// Find the top left and bottom right coordinates of a rectangle of 0's within a matrix of 1's. It's essentially a modified version of the finding the number of island problem where you only need to dfs to the right and down.
// Ex.
// [[ 1, 1, 1, 1],
// [ 1, 0, 0, 1],
// [ 1, 0, 0, 1],
// [ 1, 1, 1, 1]]
// Expected output: [[1,1], [2,2]]

// Follow up question: Expand it so it works for any number of rectangles. I ran out of time to code this part so get throught the first part quickly. Main part of this problem is updating how results are stored and tracking what's been seen.
// Ex.
// [[0, 1, 1, 1],
// [1, 0, 0, 1],
// [1, 0, 0, 1],
// [1, 1, 1, 1]]
// Expected output: [ [[0,0],[0,0]], [[1,1], [2,2]] ]

// https://leetcode.com/problems/number-of-islands

// Random string generator:

// For game recommendation, we will collect players information as parameters, such as geolocation, age, popular games among friends, etc. Then calculate a "weight" for each game, it means the possibility the game is being recommended. And we also want players to be able to see different games every time they visit the page, so we want to add some randomization in the algorithm.
// Think of your solution as a simplified game recommendation prototype, you're given a dictionary, where the keys are strings and values are positive integers representing the weight of the key (). Write a function that returns a string randomly based on its weight.

// Example:
// Input: dictionary = {'apple': 2, 'baby': 3, 'cat': 5}
// Output: "cat"

// Input: dictionary = {'apple': 2, 'baby': 3, 'cat': 5}
// Output: "baby"

package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	random := NewRandomGenerator(map[string]int{"apple": 2, "baby": 3, "cat": 5})
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())
	fmt.Println(random.pick())

}

// loop dictionary and place it into the Selection List and sort
type Selection struct {
	Key    string
	Weight int
}

type RandomGenerator struct {
	Selection      []Selection
	PrefixesWeight []int
}

func NewRandomGenerator(dict map[string]int) *RandomGenerator {

	if len(dict) == 0 {
		return nil
	}

	selections := []Selection{}
	for key, value := range dict {
		selections = append(selections, Selection{key, value})
	}

	sort.Slice(selections, func(i, j int) bool {
		return selections[i].Weight < selections[j].Weight
	})

	prefixes := make([]int, len(selections))
	prefixes[0] = selections[0].Weight
	for i := 1; i < len(selections); i++ {
		prefixes[i] = prefixes[i-1] + selections[i].Weight
	}

	return &RandomGenerator{selections, prefixes}
}

func (r *RandomGenerator) pick() string {

	fmt.Println(r.PrefixesWeight)
	fmt.Println(r.Selection)

	random := rand.Intn(r.PrefixesWeight[len(r.PrefixesWeight)-1])
	fmt.Println(random)
	choice := -1
	for i, p := range r.PrefixesWeight {
		if random < p {
			choice = i
			break
		}
	}

	return r.Selection[choice].Key
}
