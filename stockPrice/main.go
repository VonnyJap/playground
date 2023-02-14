package main

import (
	"fmt"
	"sort"
)

func main() {
	stock := Constructor()
	(&stock).Update(1, 10)
	(&stock).Update(2, 5)

	fmt.Println((&stock).Current())
	fmt.Println((&stock).Maximum())

	(&stock).Update(1, 3)

	fmt.Println((&stock).Maximum())
	(&stock).Update(4, 2)
	fmt.Println((&stock).Minimum())

}

type StockPrice struct {
	IndexedItems map[int]int
	LatestTime   int
	Prices       []int
}

// sort according to timestamp each time there is an update
// first check if there is a same timestamp corresponding to it

func Constructor() StockPrice {
	return StockPrice{
		IndexedItems: map[int]int{},
	}
}

func (this *StockPrice) Update(timestamp int, price int) {

	if timestamp > this.LatestTime {
		this.LatestTime = timestamp
	}
	this.IndexedItems[timestamp] = price
	this.Prices = []int{}
	for _, p := range this.IndexedItems {
		this.Prices = append(this.Prices, p)
	}
	sort.Ints(this.Prices)
}

func (this *StockPrice) Current() int {
	return this.IndexedItems[this.LatestTime]
}

func (this *StockPrice) Maximum() int {
	return this.Prices[len(this.Prices)-1]
}

func (this *StockPrice) Minimum() int {
	return this.Prices[0]
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
