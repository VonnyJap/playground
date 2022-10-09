package main

import "fmt"

func main() {
	ctr := Constructor()
	(&ctr).Hit(1)
	(&ctr).Hit(2)
	fmt.Println((&ctr).GetHits(3))
}

type HitCounter struct {
	timestamps []int
}

func Constructor() HitCounter {
	return HitCounter{timestamps: []int{}}
}

func (this *HitCounter) Hit(timestamp int) {
	this.timestamps = append(this.timestamps, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {

	for len(this.timestamps) > 0 {
		var diff = timestamp - this.timestamps[0]
		if diff < 300 {
			break
		}
		this.timestamps = this.timestamps[1:]

	}
	return len(this.timestamps)
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */
