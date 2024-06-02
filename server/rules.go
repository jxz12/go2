package main

import (
	"bytes"
	"strconv"
)

type Board [][]int

func NewBoard(size int) Board {
	board := make(Board, size)
	for i := 0; i < size; i++ {
		board[i] = make([]int, size)
	}
	return board
}

func (board Board) Width() int {
	return len(board)
}

func (board Board) Get(row int, col int) int {
	return board[row][col]
}

func (board Board) String() string {
	buffer := bytes.NewBufferString("")

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			buffer.WriteString(strconv.Itoa(board[row][col]))
			if col < len(board[row])-1 {
				buffer.WriteString(" ")
			}
		}
		if row < len(board)-1 {
			buffer.WriteString("\n")
		}
	}
	return buffer.String()
}

func (board Board) Score() map[int]int {
	// we use Area scoring (https://en.wikipedia.org/wiki/Rules_of_Go#Area_scoring)
	//   because Territory scoring requires saving state of captured stones
	//   plus they almost always lead to the same result

	// Algorithm is:
	//   iterate through all intersections one by one
	//     if a square is occupied, then try to capture empty components in all four directions
	//     if capturable, then pseudo-capture that component by placing a negative number
	//   then can just sum the number of stones and pseudo-captures for each player
	//   note that this means a single stone on the board means it owns the entire board hahaha

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] > 0 {
				capturer := board[row][col]
				if board.CanCapture(capturer, row-1, col, false) {
					board.Capture(-capturer, row-1, col)
				}
				if board.CanCapture(capturer, row+1, col, false) {
					board.Capture(-capturer, row+1, col)
				}
				if board.CanCapture(capturer, row, col-1, false) {
					board.Capture(-capturer, row, col-1)
				}
				if board.CanCapture(capturer, row, col+1, false) {
					board.Capture(-capturer, row, col+1)
				}
			}
		}
	}

	// TODO: at this point the board shows an explanation of the score using negative numbers
	//       it would be better to return this to the UI instead of just the score

	scores := make(map[int]int)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if board[row][col] < 0 {
				scores[-board[row][col]] += 1
				board[row][col] = -board[row][col]
			} else if board[row][col] > 0 {
				scores[board[row][col]] += 1
			}
		}
	}
	return scores
}

func (board Board) Play(player int, row int, col int) bool {
	if player == 0 {
		return false // 0 cannot be a player id
	}
	// Rule 7 (https://en.wikipedia.org/wiki/Rules_of_Go#Moving)
	// Step 1: intersection must be empty
	if board[row][col] != 0 {
		return false
	}
	board[row][col] = player

	// Step 2: look for groups of stones with liberties gone
	//   requires two passes since to-be-captured stones may still capture other stones
	//   1. determine if connected components have no liberties
	//   2. remove all stones in components with no liberties
	up := board.CanCapture(player, row-1, col, true)
	down := board.CanCapture(player, row+1, col, true)
	left := board.CanCapture(player, row, col-1, true)
	right := board.CanCapture(player, row, col+1, true)

	if up {
		board.Capture(0, row-1, col)
	}
	if down {
		board.Capture(0, row+1, col)
	}
	if left {
		board.Capture(0, row, col-1)
	}
	if right {
		board.Capture(0, row, col+1)
	}

	// Step 3: would be to remove own coloured stones, but we do not allow this due to
	// Optional Rule 7A (https://en.wikipedia.org/wiki/Rules_of_Go#Self-capture)
	if !(up || down || left || right) {
		if board.CanCapture(0, row, col, true) {
			board[row][col] = 0
			return false
		}
	}
	// Rule 8 (https://en.wikipedia.org/wiki/Rules_of_Go#Ko)
	//   is to prevent never-ending games, but we don't have to worry about it
	//   because the asynchronous nature means other better moves are available
	//   it is also more elegant to not have to remember previous board states
	return true
}
func (board Board) CanCapture(capturer int, row int, col int, prisonersNotTerritory bool) bool {
	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
		return false
	}
	if board[row][col] == capturer || board[row][col] == -capturer {
		// prevents infinite recursion if already captured or pseudo-captured
		return false
	}
	captured := board[row][col]

	var DFS func(int, int) bool
	DFS = func(row int, col int) bool {
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
			return true
		}
		// I expect/hope that this if statement is optimised away in compile time
		if prisonersNotTerritory {
			if board[row][col] == 0 {
				// empty square means liberty is available
				return false
			}
			if board[row][col] != captured {
				// captured by another square regardless of colour
				return true
			}
		} else {
			if board[row][col] == capturer || board[row][col] == -1 {
				// captured by capturer or already explored
				return true
			} else if board[row][col] != captured {
				// territory is non-collaborative so seeing any other player breaks the capture
				return false
			}
		}
		// mark intersection as explored by setting to negative
		board[row][col] = -1
		capturable := DFS(row-1, col) && DFS(row+1, col) && DFS(row, col-1) && DFS(row, col+1)
		board[row][col] = captured
		return capturable
	}
	return DFS(row, col)
}

func (board Board) Capture(replacement int, row int, col int) {
	captured := board[row][col]
	if captured == replacement {
		// prevents infinite recursion when another Capture removes the piece that was here
		return
	}

	var DFS func(int, int)
	DFS = func(row int, col int) {
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
			return
		}
		if board[row][col] != captured {
			return
		}
		board[row][col] = replacement
		DFS(row-1, col)
		DFS(row+1, col)
		DFS(row, col-1)
		DFS(row, col+1)
	}
	DFS(row, col)
}
