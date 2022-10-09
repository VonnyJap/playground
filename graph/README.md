### Graph
- Vertices
- Edges

### Graph Basic Operations
- Add Vertex
- Add Edge
- Display Vertex

### DFS
- Using stack
- keep track on the stack and visited set
- tree traversal is actually DFS

##### Algorithm DFS(G, v)
```
if v is already visited
  return        
Mark v as visited.
// Perform some operation on v.
for all neighbors x of v
  DFS(G, x)
```

### BFS
- Using queue
- keep track on the queue and visited set

##### Algorithm BFS(G, v)
```
Q = new empty FIFO queue
Mark v as visited.
Q.enqueue(v)
while Q is not empty
  a = Q.dequeue()
  // Perform some operation on a.
  for all unvisited neighbors x of a
    Mark x as visited.
    Q.enqueue(x)
```