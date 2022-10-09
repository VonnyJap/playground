#### Properties
- Simple base case: terminating scenario
- Recurrence relation: reduces all other cases towards the base case, the relationship between the result of a problem and the result of its subproblems

#### Guidelines
- we have a function f(X)
- X = [xo,x1,x2]
- call the function recursively to solve subproblems of X
- process the recursive results to solve X

#### Resources
- https://leetcode.com/explore/featured/card/recursion-i

#### Time complexity
- most of the time it will be exponential (2^n)
- example fibo number creates a tree which takes complexity of 2^n
- but can be reduced to linear with memoization

#### Space complexity
- recursive related space: stored in stack
  - returning address to the function call
  - parameters passed to the function
  - local variables within function 
- stack overflow: where the provided stack is running out
- non-recursion related space: normally stored in heap
  - example: global var
  - intermediate results


#### Tail recursion
- recursion where the recursive call is the final instruction in the recursion function
- And there should be only one recursive call in the function
- benefit: reduce stack overhead