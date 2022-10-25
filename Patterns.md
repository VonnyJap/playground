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


#### Top K elements


#### K-way Merge

#### Topological sort







