package main

import (
	"fmt"
	"sort"
)

type TopVotedCandidate struct {
	persons []int
	times   []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {

	leaders, votes := make([]int, len(persons)), make([]int, len(persons))
	leader := persons[0]
	for i := 0; i < len(persons); i++ {
		p := persons[i]
		votes[p]++
		if votes[p] >= votes[leader] {
			leader = p
		}
		leaders[i] = leader
	}
	fmt.Println(votes)
	fmt.Println(leaders)
	return TopVotedCandidate{persons: leaders, times: times}

}
func (tvc *TopVotedCandidate) Q(t int) int {
	i := sort.Search(len(tvc.times), func(p int) bool { return tvc.times[p] > t })
	fmt.Println("i: ", i)
	return tvc.persons[i-1]
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

// func (this *TopVotedCandidate) Q(t int) int {

// 	// given a Q - i will check which index that I need to stop
// 	var index = 0
// 	for i, v := range this.times {
// 		if v > t {
// 			break
// 		}
// 		index = i
// 	}
// 	fmt.Println(index)

// 	fmt.Println(this.persons[0 : index+1])

// 	// create a map to loop through the index
// 	// var dict map[int]int

// 	// and then loop through a map to find the thing or it can be better
// 	return 0
// }

func main() {
	tvc := Constructor([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})
	fmt.Println((&tvc).Q(3))
	fmt.Println((&tvc).Q(12))
	fmt.Println((&tvc).Q(25))
	fmt.Println((&tvc).Q(15))
	fmt.Println((&tvc).Q(24))
	fmt.Println((&tvc).Q(8))
	GuessingGame()
}

// Input
// ["TopVotedCandidate", "q", "q", "q", "q", "q", "q"]
// [[[0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]], [3], [12], [25], [15], [24], [8]]
// Output
// [null, 0, 1, 1, 0, 0, 1]

// Explanation
// TopVotedCandidate topVotedCandidate = new TopVotedCandidate([0, 1, 1, 0, 0, 1, 0], [0, 5, 10, 15, 20, 25, 30]);
// topVotedCandidate.q(3); // return 0, At time 3, the votes are [0], and 0 is leading.
// topVotedCandidate.q(12); // return 1, At time 12, the votes are [0,1,1], and 1 is leading.
// topVotedCandidate.q(25); // return 1, At time 25, the votes are [0,1,1,0,0,1], and 1 is leading (as ties go to the most recent vote.)
// topVotedCandidate.q(15); // return 0
// topVotedCandidate.q(24); // return 0
// topVotedCandidate.q(8); // return 1
