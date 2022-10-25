package main

type TicTacToe struct {
}

func Constructor(n int) TicTacToe {

}

func (this *TicTacToe) Move(row int, col int, player int) int {

}


["TicTacToe", "move", "move", "move", "move", "move", "move", "move"]
[[3], [0, 0, 1], [0, 2, 2], [2, 2, 1], [1, 1, 2], [2, 0, 1], [1, 0, 2], [2, 1, 1]]

// init a grid
// for each player - has a map to hold?

// each time a new move is added
// check if this moves result in a player winning?

// once win - no more moves - so stop
// if move is still allowed nobody wins

// |1||2|
// ||||
// |||1|

// class TicTacToe(object):

//     def __init__(self, n):
//         self.row, self.col, self.diag, self.anti_diag, self.n = [0] * n, [0] * n, 0, 0, n
        
//     def move(self, row, col, player):
//         offset = player * 2 - 3
//         self.row[row] += offset
//         self.col[col] += offset
//         if row == col:
//             self.diag += offset
//         if row + col == self.n - 1:
//             self.anti_diag += offset
//         if self.n in [self.row[row], self.col[col], self.diag, self.anti_diag]:
//             return 2
//         if -self.n in [self.row[row], self.col[col], self.diag, self.anti_diag]:
//             return 1
//         return 0