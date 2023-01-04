package main

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool { return true }

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int { return 0 }

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger { return []*NestedInteger{} }

type NestedIterator struct {
	NestedList        []*NestedInteger
	NextNestedInteger *NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	if len(nestedList) == 0 {
		return &NestedIterator{
			NestedList: nestedList,
		}
	}
	return &NestedIterator{
		NestedList:        nestedList[1:],
		NextNestedInteger: nestedList[0],
	}
}

func (this *NestedIterator) Next() int {

	next := this.NextNestedInteger
	if len(this.NestedList) == 0 {
		this.NextNestedInteger = nil
	} else {
		this.NextNestedInteger = this.NestedList[0]
		this.NestedList = this.NestedList[1:]
	}
	if next.IsInteger() {
		return next.GetInteger()
	}

	nextList := next.GetList()
	// what will happen to the nested thing?
}

// how do I know if NestedInteger HasNext?
func (this *NestedIterator) HasNext() bool {
	if this.NextNestedInteger == nil {
		return false
	}
	return true
}
