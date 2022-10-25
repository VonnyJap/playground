package main

import "fmt"

// for the lru
// each time get - I need to search the key and put it behind
// so each time we will take out from the beginning of stack
func main() {
	cache := Constructor(2)
	fmt.Println((&cache).Get(2))
	(&cache).Put(2, 6)
	fmt.Println((&cache).Get(1))
	(&cache).Put(1, 5)
	(&cache).Put(1, 2)
	fmt.Println((&cache).Get(1))
	fmt.Println((&cache).Get(2))
}

type LRUCache struct {
	Content  map[int]int
	History  []int
	Capacity int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Capacity: capacity,
		Content:  make(map[int]int, capacity),
		History:  []int{},
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.Content[key]; !ok {
		return -1
	}
	this.UpdateHistory(key)
	return this.Content[key]
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Content[key]; !ok {
		if this.Capacity == len(this.Content) {
			delete(this.Content, this.History[0])
			this.History = this.History[1:]
		}
	}
	this.UpdateHistory(key)
	this.Content[key] = value
}

func (this *LRUCache) UpdateHistory(key int) {

	for i, h := range this.History {
		if h == key {
			if i == len(this.History)-1 {
				this.History = append(this.History[:i])
			} else {
				this.History = append(this.History[:i], this.History[i+1:]...)
			}
			break
		}
	}
	this.History = append(this.History, key)
}
