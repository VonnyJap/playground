#### Sliding Window
- Maximum sum subarray of size "K" (easy)
- Longest substring with "K" distinct characters (medium)
- What is this?
  - assign the first pointer to one index signifying where the window starts, and the second pointer to another index signifying where the window ends.
  - between these two points, gives a subset of data
  - shift to get new subset

[Sliding Window Algorithm Approach(Go example)](https://sdever.medium.com/sliding-window-algorithm-approach-go-example-9df848c76077)


#### Two pointers
- Reverse arr
- Rotate arr
- Squaring a sorted array

- Implementation for rotating array
  - reverse the arr using the two pointers
  - reverse arr[:len(nums)-k] => a
  - reverse arr[len(nums)-k:] => b
  - c => append(a,b)
  - reverse c

#### Fast and Slow pointers
- Linked list or array
- Linked list has cycle

#### Merge Intervals
[Golang Implementation](https://golangbyexample.com/merge-overlapping-intervals-golang/)

#### Cyclic sort


#### In-place reversal of linked list


#### Tree BFS
- If you’re asked to traverse a tree in a level-by-level fashion (or level order traversal)
- Using queue
- Level by Level to keep tracking on the leveling
- Binary Tree Level Order Traversal (easy)


#### Tree DFS
- We can use stack or recursion
- If you’re asked to traverse a tree with in-order, preorder, or postorder DFS
- If the problem requires searching for something where the node is closer to a leaf

#### Two heaps

#### Subsets


#### Modified binary search
- Given a sorted arr, linked list, matrix
- Steps:
  - find the middle of start and end
  - if target same with middle then return middle
  - if target larger than middle -> start = middle + 1
  - else -> start = middle - 1

#### Top K elements
- The best data structure to keep track of "K" elements is Heap
- Steps:
  - Insert "K" elements into the min-heap or max-heap based on the problem.
  - Iterate through the remaining numbers and if you find one that is larger than what you have in the heap, then remove that number and insert the larger one.

- Implement heap in Golang
```
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}
```

- Practical heap example: find largest k element
```
func FindBestKElements(nums []int, k int) []int {
	h := &IntHeap{}
	for _, val := range nums { // O(N)
		heap.Push(h, val) // O(log K)
		if h.Len() > k {
			heap.Pop(h) // O(log K)
		}
	}

	return func() []int { // O (k log k)
		result := make([]int, h.Len())
		initialLen := h.Len()
		for i := initialLen; i > 0; i-- {
			result[i-1] = heap.Pop(h).(int)
		}
		return result
	}()
}
```

#### K-way Merge

#### Topological sort

#### Union Find
- detect cycle in undirected graph
- [Algo explained](https://www.geeksforgeeks.org/union-find/)

#### Greedy
- This is more like a technique like DP, Divide and Conquer, Brute Force
- If an optimal solution to the problem can be found by choosing the best choice at each step without reconsidering the previous steps once chosen, the problem can be solved using a greedy approach. This property is called greedy choice property.
- If the optimal overall solution to the problem corresponds to the optimal solution to its subproblems, then the problem can be solved using a greedy approach. This property is called optimal substructure.
- Steps:
  - To begin with, the solution set (containing answers) is empty.
  - At each step, an item is added to the solution set until a solution is reached.
  - If the solution set is feasible, the current item is kept.
  - Else, the item is rejected and never considered again.
- Examples: largest sum, shortest path, Minimum Spanning Trees
- Does not work for path with negative value
- References
  - [Programiz](https://www.programiz.com/dsa/greedy-algorithm)
  - [Simplilearn](https://www.simplilearn.com/tutorials/data-structure-tutorial/greedy-algorithm)
  - [Top Questions](https://medium.com/techie-delight/top-7-greedy-algorithm-problems-3885feaf9430)

#### Modified binary search to find closest value?

#### Recursive with Memoization: top down DP

#### Dynamic Programming
- solving a problem using subproblems and storing intermediate solutions, saves computation time at the expense of storage space
- top down: store intermediate result in array or hash table, modified recursion
- bottom up: iterative version of DP, the idea is the solution of a subproblem depends on the solution of smaller subproblems

```
// memoized version
static int fibonacciMemo(int n) {
    // entry table to cache the computed values
    int[] fibs = new int[n + 1];
    // initialize entry table with -1 to say value has to calculated
    Arrays.fill(fibs, -1);

    return fib(n, fibs);
}

static int fib(int n, int[] fibs) {
    if (n == 0 || n == 1) return n;

    if (fibs[n] == -1) {
        fibs[n] = fib(n - 1, fibs) + fib(n - 2, fibs);
    }

    return fibs[n];
}
```

- The best way to practice dynamic programming is:
  1. First, define a brute force recursive solution.
  2. Characterise the structure of the recursive solution.
  3. Identify the base cases.
  4. Store the computed values of overlapping subproblems.
  5. Convert Recursive code to Memoised code.
  6. Convert Memoised code to Tabular form.
- Normally memoized approach is solved by storing the target value because there will be multiple calls for it
- and the next time the call is happening, we can just use that value before proceeding to the logic further


#### Double Linked List
- [LRU Cache](https://golangbyexample.com/lru-cache-implementation-golang/)

#### arithmetic series sum
- formula i.e. ( n * (n + 1) ) / 2.

#### find intersection of two lists
- use two pointers approach
- hashmap

#### Bits manipulation
- & is a bitwise AND operator. Doing (num & 1) checks if the last bit (least significant bit) of the number is set. If it is set, the number is odd, and if it is not set, it is even.
- AND: extract a subset of bits in a value
- OR: set a subset of bits in a value
- XOR: toggle a subset of bits in a value
  - 1 ^ 1 = 0
  - 0 ^ 1 = 1
  - 1 ^ 0 = 1
  - 0 ^ 0 = 0
- complement in Go can be done using 1 ^ M

#### Reference
[14 Patterns to Ace Any Coding Interview Question](https://hackernoon.com/14-patterns-to-ace-any-coding-interview-question-c5bb3357f6ed)
[Grokking-the-Coding-Interview-Patterns-for-Coding-Questions](https://github.com/cl2333/Grokking-the-Coding-Interview-Patterns-for-Coding-Questions)

[TREE QUESTION PATTERN 2023 || TREE STUDY GUIDE](https://leetcode.com/discuss/study-guide/2879240/tree-question-pattern-2023-tree-study-guide)
[Graph Algorithm](https://leetcode.com/discuss/interview-question/2886740/familiar-with-graphs-algorithm)

#### Airbnb
[Question about iterator](https://leetcode.com/discuss/interview-question/352704/airbnb-phone-screen-removing-list-iterator)

#### Backtracking Algo
- uses a brute force approach for finding the desired output
- if current solution not suitable then backtrack to try other solutions
- recursive is used
- type:
  - decision problem: finding feasible solution
  - optimization problem: finding best solution
  - enumeration problem: finding all feasible solutions
```
Backtrack(x)
    if x is not a solution
        return false
    if x is a new solution
        add to list of solutions
    backtrack(expand x)
```
[Sample Problem: Rat in a maze](https://www.geeksforgeeks.org/rat-in-a-maze-backtracking-2/)

#### BFS vs DFS
- The main difference:
  - BFS:  Tries to explore all the neighbors it can reach from the current node. It will use a queue data structure.
  - DFS:  Tries to reach the farthest node from the current node and come back (backtrack) to the current node to explore its other neighbors. This will use a stack data structure.
- We will prefer to use BFS when we know that our solution might lie closer to the starting point or if the graph has greater depths.
- We will prefer to use DFS when we know our solution might lie farthest from the starting point or when the graph has a greater width.
- If we have multiple starting points and the problem requires us to start traversing all those starting points parallelly then we can think of BFS as we can push all those starting points in the queue and start exploring them first.
- It’s generally a good idea to use BFS if we need to find the shortest distance from a node in the unweighted graph.
- We will be using DFS mostly in path-finding algorithms to find paths between nodes.



#### Shortest Path
[Ref](https://pencilprogrammer.com/algorithms/shortest-path-in-unweighted-undirected-graph-using-dfs/)


#### Cyclic Sort Programming Pattern
[Ref](https://stephenjoel2k.medium.com/cyclic-sort-programming-pattern-16eb16ac9c4b)

#### K-way Merge
- How to identify:
  - The problem will feature sorted arrays, lists, or a matrix
  - If the problem asks you to merge sorted lists, find the smallest element in a sorted list
- Example:
  - Merge K sorted lists
- Steps:
  1. Insert first number from each array into the heap
  2. Pop out and get it merged into the return list
  3. Heap size = number of arrays to ensure that in each run all arrays will get turn