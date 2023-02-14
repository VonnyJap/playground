package main

func main() {}

type LRUCache struct {
	Capacity int
	Contents map[int]*DLL
	Head     *DLL
	Tail     *DLL
}

type DLL struct {
	Key   int
	Value int
	Next  *DLL
	Prev  *DLL
}

func (c LRUCache) AddNode(node *DLL) {
	node.Prev = c.Head
	node.Next = c.Head.Next

	c.Head.Next.Prev = node
	c.Head.Next = node
}

func (c LRUCache) RemoveNode(node *DLL) {

	prev := node.Prev
	next := node.Next

	prev.Next = next
	next.Prev = prev
}

func (c LRUCache) MoveToHead(node *DLL) {
	c.RemoveNode(node)
	c.AddNode(node)
}

func (c LRUCache) Pop() *DLL {

	prev := c.Tail.Prev
	c.RemoveNode(prev)

	return prev
}

func Constructor(capacity int) LRUCache {

	head := &DLL{}
	tail := &DLL{}

	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Capacity: capacity,
		Contents: map[int]*DLL{},
		Head:     head,
		Tail:     tail,
	}

}

// to implement this we add to Contents

func (this *LRUCache) Get(key int) int {
	if dll, ok := this.Contents[key]; ok {
		this.MoveToHead(dll)
		return dll.Value
	}
	return -1
}

// need to check if Contents == Capacity
// if yes - we need to evict
func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.Contents[key]; ok {
		this.Contents[key].Value = value
		this.MoveToHead(this.Contents[key])
		return
	}

	newNode := &DLL{Key: key, Value: value}
	this.AddNode(newNode)
	this.Contents[key] = newNode

	if len(this.Contents) > this.Capacity {
		tail := this.Pop()
		delete(this.Contents, tail.Key)
	}
}
