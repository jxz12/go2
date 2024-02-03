package main

import "bytes"
import "strconv"

type IBoard interface {
	Width() int
	Get(int, int) int
	ToString() string
	Score() map[int]int
	Play(int, int, int) bool
}
type Board [][]int

func NewBoard(size int) IBoard {
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

func (board Board) ToString() string {
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
	scores := make(map[int]int)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			scores[board[row][col]] += 1
		}
	}
	return scores
}

func (board Board) Play(player int, row int, col int) bool {
	if player == 0 {
		return false // 0 cannot be a player id
	}
	// Rule 7 (https://en.wikipedia.org/wiki/Rules_of_Go#Moving)
	// Step 1: position must be empty
	if board[row][col] != 0 {
		return false
	}
	board[row][col] = player

	// Step 2: look for groups of stones with liberties gone
	//   requires two passes since to-be-captured stones may still capture other stones
	//   1. determine if connected components have no liberties
	//   2. remove all stones in components with no liberties
	up := board.CanCapture(player, row-1, col)
	down := board.CanCapture(player, row+1, col)
	left := board.CanCapture(player, row, col-1)
	right := board.CanCapture(player, row, col+1)

	if up {
		board.Capture(row-1, col)
	}
	if down {
		board.Capture(row+1, col)
	}
	if left {
		board.Capture(row, col-1)
	}
	if right {
		board.Capture(row, col+1)
	}

	// Step 3: would be to remove own coloured stones, but we do not allow this due to
	// Optional Rule 7A (https://en.wikipedia.org/wiki/Rules_of_Go#Self-capture)
	if !(up || down || left || right) {
		if board.CanCapture(0, row, col) {
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
func (board Board) CanCapture(capturer int, row int, col int) bool {
	if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
		return false
	}
	if board[row][col] == capturer {
		return false
	}
	captured := board[row][col]

	var DFS func(int, int) bool
	DFS = func(row int, col int) bool {
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
			return true
		}
		if board[row][col] == 0 {
			// empty square means liberty is available
			return false
		}
		if board[row][col] != captured {
			// liberty taken by another colour stone
			return true
		}
		// mark position as explored by setting to negative number
		board[row][col] = -board[row][col]
		capturable := DFS(row-1, col) && DFS(row+1, col) && DFS(row, col-1) && DFS(row, col+1)
		board[row][col] = -board[row][col]
		return capturable
	}
	return DFS(row, col)
}

func (board Board) Capture(row int, col int) {
	player := board[row][col]
	if player == 0 {
		// prevents infinite recursion when another Capture removes the piece that was here
		return
	}

	var DFS func(int, int)
	DFS = func(row int, col int) {
		if row < 0 || row > len(board)-1 || col < 0 || col > len(board[row])-1 {
			return
		}
		if board[row][col] != player {
			return
		}
		board[row][col] = 0
		DFS(row-1, col)
		DFS(row+1, col)
		DFS(row, col-1)
		DFS(row, col+1)
	}
	DFS(row, col)
}
