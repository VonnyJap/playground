package main

// n： 1, 2, 3, 4, 5, 6, 7
// F(n): 1, 2, 2, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 6, 7, 7, 7, 7
// Given n, return the nth number in F(n)

// # my naive solution
// def repeat(n):
//     if n == 1:
//         return [1]
//     if n == 2:
//         return [1, 2, 2]

//     F = [1, 2, 2]

//     for i in range(3, n+1):
//         F += [i] * F[i-1]

//     print(F, len(F))
//     return F[n]
// Follow up: If repeat() is called very frequently, how to optimize the performance? if n is very large, how to reduce memory usage?
