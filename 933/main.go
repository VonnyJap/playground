package main

import "fmt"

func main() {
	ctr := Constructor()
	fmt.Println(ctr.Ping(1))
	fmt.Println(ctr.Ping(100))
	fmt.Println(ctr.Ping(3001))
	fmt.Println(ctr.Ping(3002))
}

type RecentCounter struct {
	Calls []int
}

func Constructor() RecentCounter {
	return RecentCounter{Calls: []int{}}
}

func (this *RecentCounter) Ping(t int) int {

	// var valid = []int{}

	// for _, c := range this.Calls {
	// 	if c >= t-3000 {
	// 		valid = append(valid, c)
	// 	}
	// }

	for i, rlen := 0, len(this.Calls); i < rlen; i++ {
		j := i - (rlen - len(this.Calls))
		if this.Calls[j] < t-3000 {
			this.Calls = append(this.Calls[:j], this.Calls[j+1:]...)
		}
	}

	this.Calls = append(this.Calls, t)
	return len(this.Calls)
}
