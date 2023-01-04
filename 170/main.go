package main

import (
	"fmt"
	"sort"
)

func main() {

	twoSum := Constructor()

	(&twoSum).Add(1)
	(&twoSum).Add(3)
	(&twoSum).Add(5)
	fmt.Println((&twoSum).Find(4))
	fmt.Println((&twoSum).Find(7))

}

type TwoSum struct {
	Members []int
}

func Constructor() TwoSum {
	return TwoSum{Members: []int{}}
}

func (this *TwoSum) Add(number int) {
	this.Members = append(this.Members, number)
	sort.Ints(this.Members)
}

func (this *TwoSum) Find(value int) bool {

	start := 0
	end := len(this.Members) - 1

	for start < end {
		sum := this.Members[start] + this.Members[end]
		if sum == value {
			return true
		}
		if sum < value {
			start++
			continue
		}
		end--
	}
	return false
}

/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
